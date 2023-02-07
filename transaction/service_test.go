package transaction_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/ayushjnv1/Gobank/db"
	mockDb "github.com/ayushjnv1/Gobank/db/mocks"
	"github.com/ayushjnv1/Gobank/transaction"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestSuiteService struct {
	suite.Suite
	store       *mockDb.Storer
	transection transaction.Service
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuiteService))
}

func (t *TestSuiteService) SetupTest() {
	t.store = &mockDb.Storer{}
	t.transection = transaction.NewTransactionService(t.store)
}

func (suite *TestSuiteService) TestAmounttransection() {
	t := suite.T()
	ctx := context.Background()
	accountTest1 := db.Account{
		UserID:  "1",
		Balance: 1200,
		ID:      "1234edfdsx",
	}
	accountTest2 := db.Account{
		UserID:  "12",
		Balance: 1300,
		ID:      "1234edfdsx"}
	transactionObj := transaction.TransactionRequest{
		Amount:    1200,
		CreditAcc: "1234edf",
		DebitAcc:  "1234edfdsx",
	}
	transactionDBObj := db.TransactionStruct{
		Amount: 1200,
		CreditAcc: sql.NullString{
			String: "1234edf",
			Valid:  true,
		},
		DebitAcc: sql.NullString{
			String: "1234edfdsx",
			Valid:  true,
		},
	}
	t.Run("Successful transecetion ", func(t *testing.T) {
		suite.SetupTest()

		suite.store.On("GetAccountBalance", ctx, accountTest1.ID).Once().Return(1200, nil)
		suite.store.On("GetAccount", ctx, accountTest1.ID).Return(accountTest1, nil)
		suite.store.On("Amounttransaction", ctx, transactionDBObj).Return(nil)
		suite.store.On("GetAccountBalance", ctx, accountTest1.ID).Twice().Return(0, nil)

		amm, err := suite.transection.AmountTransaction(ctx, transactionObj, "1")
		assert.NoError(t, err)
		assert.Equal(t, amm, 0)
	})
	t.Run("unsuccessful transecetion insufficient Ammount", func(t *testing.T) {
		suite.SetupTest()
		transactionDBObj.Amount = 1300
		suite.store.On("GetAccountBalance", ctx, accountTest1.ID).Once().Return(1200, nil)
		suite.store.On("GetAccount", ctx, accountTest1.ID).Return(accountTest1, nil)
		suite.store.On("Amounttransaction", ctx, transactionDBObj).Return(nil)
		suite.store.On("GetAccountBalance", ctx, accountTest1.ID).Twice().Return(0, nil)
		transactionObj.Amount = 1300
		_, err := suite.transection.AmountTransaction(ctx, transactionObj, "1")
		assert.EqualError(t, err, transaction.ErrInSufficientAmmount.Error())

	})
	t.Run("unsuccessful transecetion invalid user", func(t *testing.T) {
		suite.SetupTest()

		suite.store.On("GetAccountBalance", ctx, accountTest2.ID).Once().Return(1300, nil)
		suite.store.On("GetAccount", ctx, accountTest1.ID).Return(accountTest2, nil)
		suite.store.On("Amounttransaction", ctx, transactionDBObj).Return(nil)
		suite.store.On("GetAccountBalance", ctx, accountTest1.ID).Twice().Return(0, nil)

		_, err := suite.transection.AmountTransaction(ctx, transactionObj, "2")
		assert.EqualError(t, err, transaction.ErrUnAuthorize.Error())

	})
}
