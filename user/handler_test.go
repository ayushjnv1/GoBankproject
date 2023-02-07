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

func (suite *TestSuiteHandler) TestUpdateUser() {
	t := suite.T()

	fmt.Println("call ")
	url := "/users/{id}"
	r := httptest.NewRequest(http.MethodPut, url, strings.NewReader(`{
		"name":"ayush",
		"email":"ayush12@gmail1.com",
		"role":"admin"	
	}`))

	r = mux.SetURLVars(r, map[string]string{
		"id": "1",
	})
	w := httptest.NewRecorder()

	userObj := user.UpdateUser{
		Name:  "ayush",
		Email: "ayush12@gmail1.com",
		Role:  "admin",
	}

	t.Run("when user update success", func(t *testing.T) {
		suite.SetupTest()
		suite.service.On("UpadateUser", r.Context(), userObj, "1").Return(nil)
		user.UpdateUserById(suite.service)(w, r)
		assert.Equal(t, http.StatusOK, w.Code)
	})

}

func (suite *TestSuiteHandler) TestList() {
	t := suite.T()
	url := "/users"
	r := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	userli := []user.UserResp{{
		ID:    "1",
		Name:  "ayush",
		Email: "ayush",
		Role:  "admin"}}
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
