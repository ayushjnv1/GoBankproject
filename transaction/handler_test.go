package transaction_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ayushjnv1/Gobank/transaction"
	"github.com/ayushjnv1/Gobank/transaction/mocks"
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

func (suite *TestSuiteHandler) TestTransection() {
	t := suite.T()
	url := "/amountTransaction"
	r1 := httptest.NewRequest(http.MethodPost, url, strings.NewReader(`{
		"amount":1800,
		"credit_acc":"dummy:account",
		"debit_acc":"dummy::string"
	}`))

	r1.Header.Set("id", "1")
	w1 := httptest.NewRecorder()
	t.Run("transection succesfully", func(t *testing.T) {
		suite.SetupTest()
		suite.service.On("Amounttransection", r1.Context(), 1800, "dummy:account", "dummy::string", "1").Return(1200, nil)
		transaction.Amounttransaction(suite.service)(w1, r1)
		assert.Equal(t, 202, w1.Code)

	})

	r2 := httptest.NewRequest(http.MethodPost, url, strings.NewReader(`{
		"amount":1800,
        "credit_acc":"dummy:account",    
	}`))
	w2 := httptest.NewRecorder()
	r2.Header.Set("id", "1")
	t.Run("transection unsuccesfully", func(t *testing.T) {
		suite.SetupTest()
		suite.service.On("Amounttransection", r2.Context(), 1800, "dummy:account", "", "1").Return(1200, nil)
		transaction.Amounttransaction(suite.service)(w2, r2)
		assert.Equal(t, http.StatusBadRequest, w2.Code)
	})
	assert.Equal(t, http.StatusBadRequest, w2.Code)

}
