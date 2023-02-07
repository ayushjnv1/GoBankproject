package user_test

import (
	"context"
	"testing"

	"github.com/ayushjnv1/Gobank/db"
	mockDb "github.com/ayushjnv1/Gobank/db/mocks"

	"github.com/ayushjnv1/Gobank/user"
	"github.com/ayushjnv1/Gobank/user/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	store   *mockDb.Storer
	user    user.Service
	encrypt *mocks.Encrypt
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (t *TestSuite) SetupTest() {
	t.store = &mockDb.Storer{}
	t.user = &user.UserService{Store: t.store, Encrypt: &mocks.Encrypt{}}
	t.encrypt = &mocks.Encrypt{}
}

func (suite *TestSuite) TestUpadateUserService() {
	t := suite.T()
	ctx := context.Background()
	testUser := user.UpdateUser{Name: "Ayush", Email: "ayushjnv1@gmail.com", Role: "admin"}
	testUser1 := user.UpdateUser{Email: "ayushjnv1@gmail.com", Role: "admin"}
	testUser2 := user.UpdateUser{Name: "Ayush", Role: "admin"}
	t.Run("when user update success", func(t *testing.T) {
		suite.SetupTest()
		suite.store.On("UpdateUser", ctx, db.User{
			Name:  testUser.Name,
			Email: testUser.Email,
			Role:  testUser.Role}, "1").Return(nil)
		err := suite.user.UpadateUser(ctx, testUser, "1")
		assert.NoError(t, err)

	})
	t.Run("when user update fail name not present ", func(t *testing.T) {
		suite.SetupTest()
		suite.store.On("UpdateUser", ctx, db.User{
			Name:  testUser1.Name,
			Email: testUser1.Email,
			Role:  testUser1.Role}, "1").Return(nil)

		goterr := suite.user.UpadateUser(ctx, testUser1, "1")
		assert.Error(t, goterr, user.ErrEmptyName)

	})
	t.Run("when user update fail email not present ", func(t *testing.T) {
		suite.SetupTest()
		suite.store.On("UpdateUser", ctx, db.User{
			Name:  testUser2.Name,
			Email: testUser2.Email,
			Role:  testUser2.Role}, "1").Return(nil)

		goterr := suite.user.UpadateUser(ctx, testUser2, "1")
		assert.EqualError(t, goterr, user.ErrEmptyEmail.Error())

	})

}
