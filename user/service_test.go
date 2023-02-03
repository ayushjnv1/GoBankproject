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

type testCreacteUser struct{
	user user.UserCreate
	err error

}

type TestSuite struct{
 suite.Suite
 store *mockDb.Storer
 user user.Service
 encrypt *mocks.Encrypt
}


func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (t *TestSuite) SetupTest(){
	t.store = &mockDb.Storer{}	
	t.user = &user.UserService{Store:t.store,Encrypt: &mocks.Encrypt{}}
	t.encrypt = &mocks.Encrypt{}
}

// func (suite *TestSuite) TearDownSuite() {
// 	suite.store.AssertExpectations(suite.T())
// 	suite.encrypt.AssertExpectations(suite.T())
// }

// func (suite *TestSuite) TestCreateUser(){
// 	t := suite.T()
// 	// ctx := context.Background()
// 	// assert := assert.New(t)
// 	// db := &mocks.Storer{}
// 	// userService := NewUserService(db)
// 	p:= "1234"
// 	ancryptReal := NewEncrypt();
// 	p1,_:= ancryptReal.HashPassword("1234")
// 	fmt.Println(p1,"hash..............")
// 	testUser := []testCreacteUser{
// 		{user:UserCreate{Id:"12",Email:"ayush@gmail.com",Password:"1234",Role:"admin",Name:"ayu"},err:nil},
// 		{user:UserCreate{Id:"12",Email:"ayush@gmail.com",Role:"admin",Name:"ayu"},err:ErrEmptyPassword},
// 		{user:UserCreate{Id:"12",Email:"ayush@gmail.com",Password:p,Name:"ayu"},err:ErrEmptyRole},
// 		{user:UserCreate{Id:"12",Email:"ayush@gmail.com",Password:p,Role:"admin"},err:ErrEmptyName},
// 		{user:UserCreate{Id:"12",Password:p,Role:"admin",Name:"ayu"},err:ErrEmptyEmail},
// 	}
// 	// for _,item:=range(testUser){

// 	// 	err := userService.CreateUser(ctx,item.user)
// 	// 	db.On("CreateUser",ctx,mocks.Anything)
		
// 	// }
// 	ctx := context.Background()
    
// 	item := testUser[0]
// 		suite.SetTestSuit()
		
// 		t.Run("when user creation success",func(t *testing.T) {		
// 		suite.encrypt.On("HashPassword",item.user.Password).Return(p1,nil)	
// 		suite.store.On("CreateUser",ctx,db.User{
// 			Name:item.user.Name,
// 			Email: item.user.Email,
// 			Role: item.user.Role,
// 			Password: p1,
// 		  }).Return(nil)
    
//         err:= suite.user.CreateUser(ctx,item.user)
// 		fmt.Println(err)
// 		assert.NoError(t,err)
// 		suite.TearDownSuite()
// 	})


	


// }


func (suite *TestSuite)TestUpadateUser(){
	t := suite.T()
	ctx := context.Background()
	testUser := user.UpdateUser{Name: "Ayush",Email:"ayushjnv1@gmail.com",Role:"admin" }
	testUser1 := user.UpdateUser{Email:"ayushjnv1@gmail.com",Role:"admin" }
	testUser2 := user.UpdateUser{Name: "Ayush",Role:"admin" }
	t.Run("when user update success",func(t *testing.T) {
		suite.SetupTest()		
        suite.store.On("UpdateUser",ctx,db.User{
		Name:testUser.Name,
		Email: testUser.Email,
		Role: testUser.Role},"1").Return(nil)
		err:= suite.user.UpadateUser(ctx,testUser,"1")
		assert.NoError(t,err)
			
	})
	t.Run("when user update fail name not present ",func(t *testing.T) {
		suite.SetupTest()
		suite.store.On("UpdateUser",ctx,db.User{
			Name:testUser1.Name,
			Email: testUser1.Email,
			Role: testUser1.Role},"1").Return(nil)
    
		goterr:= suite.user.UpadateUser(ctx,testUser1,"1")
		assert.Error(t,goterr,user.ErrEmptyName)
			
	})
	t.Run("when user update fail email not present ",func(t *testing.T) {
		suite.SetupTest()
		suite.store.On("UpdateUser",ctx,db.User{
			Name:testUser2.Name,
			Email: testUser2.Email,
			Role: testUser2.Role},"1").Return(nil)
    
		goterr:= suite.user.UpadateUser(ctx,testUser2,"1")		
		assert.EqualError(t,goterr,user.ErrEmptyEmail.Error())
		
	})
	

}