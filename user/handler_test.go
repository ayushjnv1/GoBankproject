package user_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ayushjnv1/Gobank/user"
	"github.com/ayushjnv1/Gobank/user/mocks"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestSuiteHandler struct {
	suite.Suite
	service *mocks.Service
}

// start point or we say tell testify run suite
func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuiteHandler))
}

func (suite *TestSuiteHandler) SetupTest() {
	suite.service = &mocks.Service{}
}

func (suite *TestSuiteHandler) TestUpdateTicket() {
	t := suite.T()

	fmt.Println("call ")
	url := "/users/{id}"
	r := httptest.NewRequest(http.MethodPut, url, strings.NewReader(`{
		"name":"ayush",
		"email":"ayush12@gmail1.com",
		"role":"admin"	
	}`))
	r.Header.Set("Content-Type", "json")
	r.Header.Set("User-Agent", "PostmanRuntime/7.29.2")
	r.Header.Set("Accept", "*/*")
	r.Header.Set("Connection", "keep-alive")
	// vars:=make(map[string]string)
	// vars["id"]="1"
	// // mux.Vars(r)["id"] = "1"
	// mux.SetURLVars(r,vars)
	// ctx := r.Context()
	// 	mux.SetURLVars(r, 0, 1)
	// ctx = context.WithValue(ctx, httprouter.ParamsKey, httprouter.Params{
	//     {"uuid", "some-uuid"},
	// })
	// r = r.WithContext(ctx)

	w := httptest.NewRecorder()

	userObj := user.UpdateUser{
		Name:  "ayush",
		Email: "ayush12@gmail1.com",
		Role:  "admin",
	}
	// mux.Vars()
	t.Run("when user update success", func(t *testing.T) {
		suite.SetupTest()
		fmt.Println("call run")
		err := suite.service.On("UpadateUser", r.Context(), userObj, "1").Return(nil)
		fmt.Println(err, "error")
		fmt.Println("this is call")
		user.UpdateUserById(suite.service)(w, r)

		fmt.Printf("%+v", w)
	})

}

func (suite *TestSuiteHandler) TestList() {
	t := suite.T()
	url := "/users"
	r := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	userli := []user.UserResp{{ID: "1", Name: "ayush", Email: "ayush", Role: "admin"}}
	userList := user.UserList{User: userli}

	t.Run("user get list", func(t *testing.T) {
		suite.SetupTest()
		suite.service.On("List", r.Context()).Return(userList, nil)
		user.ListOfUser(suite.service)(w, r)
		res, _ := json.Marshal(userList)
		expected := (string(res))
		assert.Equal(t, bytes.NewBufferString(expected), (w.Body))
		assert.Equal(suite.T(), http.StatusOK, w.Code)
	})

}
