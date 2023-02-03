package user

type UserCreate struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UpdateUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
type UserResp struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
type UserList struct {
	User []UserResp
}
type UserResponse struct {
	User UserResp
}

func ValidateCreateUser(user UserCreate) error {
	if user.Email == "" {
		return ErrEmptyEmail
	}
	if user.Password == "" {
		return ErrEmptyPassword
	}
	if user.Role == "" {
		return ErrEmptyRole
	}
	if user.Name == "" {
		return ErrEmptyName
	}
	return nil
}

func UpdateUserValidate(user UpdateUser) error {
	if user.Email == "" {
		return ErrEmptyEmail
	}

	if user.Name == "" {
		return ErrEmptyName
	}
	if user.Role == "" {
		return ErrEmptyRole
	}
	return nil
}
func IsBadRequest(err error) bool {
	return err == ErrEmptyEmail || err == ErrEmptyName || err == ErrEmptyPassword || err == ErrEmptyRole
}
