package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type User struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Password  string    `db:"password"`
	Email     string    `db:"email"`
	Role      string    `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	DeletedAt time.Time `db:"deleted_at"`
}

const (
	createUser     = "INSERT INTO user(name,password,email,role,created_at) VALUES(?,?,?,?,SYSDATE())"
	listOfUser     = "SELECT name,email,password,role,id FROM user WHERE deleted_at IS NULL"
	findById       = "SELECT name,email,password,role,id FROM user WHERE id = ? AND deleted_at IS NULL"
	updateUser     = `UPDATE user SET name=?, email=?,role=? WHERE id =?`
	deleteUser     = `UPDATE user SET deleted_at=SYSDATE() WHERE id=?`
	updatePassword = `UPDATE user SET password=? WHERE id=?`
	findbyEmail    = `SELECT id,email,password,role FROM user WHERE email=? AND deleted_at IS NULL`
)

func (s *store) CreateUser(ctx context.Context, user User) (err error) {
	fmt.Println(user, "user")
	_, err = s.db.DB.ExecContext(ctx, createUser, user.Name, user.Password, user.Email, user.Role)
	return err
}

func (s *store) DeleteUser(ctx context.Context, ID string) (err error) {

	res, err := s.db.ExecContext(ctx, deleteUser, ID)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if count == 0 {
		return ErrUserNotExist
	}
	return
}
func (s *store) UpdateUser(ctx context.Context, user User, ID string) (err error) {

	params := []interface{}{user.Name, user.Email, user.Role, ID}
	res, err := s.db.ExecContext(ctx, updateUser, params...)
	if err != nil {
		return
	}

	count, err := res.RowsAffected()
	if count == 0 {
		return errors.New("User does not exist")
	}
	return
}
func (s *store) UpdatePassword(ctx context.Context, password string, ID string) (err error) {

	params := []interface{}{password, ID}
	res, err := s.db.ExecContext(ctx, updatePassword, params...)
	if err != nil {
		return
	}

	count, err := res.RowsAffected()
	if count == 0 {
		return errors.New("User does not exist")
	}
	return
}

func (s *store) ListOfUser(ctx context.Context) (user []User, err error) {
	err = s.db.SelectContext(ctx, &user, listOfUser)
	if err == sql.ErrNoRows {
		return user, ErrUserNotExist
	}
	return
}
func (s *store) FindById(ctx context.Context, ID string) (user User, err error) {
	use := []User{}
	err = s.db.SelectContext(ctx, &use, findById, ID)

	if len(use) == 0 {
		return user, ErrUserNotExist
	}
	return use[0], err
}

func (s *store) FindByEmail(ctx context.Context, email string) (user User, err error) {
	userList := []User{}
	err = s.db.SelectContext(ctx, &userList, findbyEmail, email)

	if err != nil {
		return
	}
	if len(userList) == 0 {
		return user, ErrUserNotExist
	}
	return userList[0], err
}
