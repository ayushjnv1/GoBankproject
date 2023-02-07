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
	testrequest1 := httptest.NewRequest(http.MethodPost, url, strings.NewReader(`{
		"amount":1800,
		"credit_acc":"dummy:account",
		"debit_acc":"dummy::string"
	}`))

	testrequest1.Header.Set("id", "1")
	testResponse1 := httptest.NewRecorder()
	transactionObj := transaction.TransactionRequest{
		Amount:    1800,
		CreditAcc: "dummy:account",
		DebitAcc:  "dummy::string",
	}
	t.Run("transection succesfully", func(t *testing.T) {
		suite.SetupTest()
		suite.service.On("AmountTransaction", testrequest1.Context(), transactionObj, "1").Return(1200, nil)
		transaction.InitiateTransaction(suite.service)(testResponse1, testrequest1)
		assert.Equal(t, 202, testResponse1.Code)

	})

	testRequest2 := httptest.NewRequest(http.MethodPost, url, strings.NewReader(`{
		"amount":1800,
        "credit_acc":"dummy:account",    
	}`))
	testResponse2 := httptest.NewRecorder()
	testRequest2.Header.Set("id", "1")
	transactionObj2 := transaction.TransactionRequest{
		Amount:    1800,
		CreditAcc: "dummy:account",
	}
	t.Run("transection unsuccesfully debit Account miss", func(t *testing.T) {
		suite.SetupTest()
		suite.service.On("AmountTransaction", testRequest2.Context(), transactionObj2, "1").Return(1200, nil)
		transaction.InitiateTransaction(suite.service)(testResponse2, testRequest2)
		assert.Equal(t, http.StatusBadRequest, testResponse2.Code)
	})
	assert.Equal(t, http.StatusBadRequest, testResponse2.Code)

}
