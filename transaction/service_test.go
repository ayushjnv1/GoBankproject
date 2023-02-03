package transaction_test

import (
	"context"
	"fmt"
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
	custd := db.Customer{Userid: "1", Balance: 1200, ID: "1234edfdsx"}
	cust2 := db.Customer{Userid: "12", Balance: 1200, ID: "1234edfdsx"}
	// custc := db.Customer{Uid: "1",Amount:200,Id:"1234edf"}
	t.Run("Successful transecetion ", func(t *testing.T) {
		suite.SetupTest()

		suite.store.On("GetAccountBalance", ctx, custd.ID).Once().Return(1200, nil)
		suite.store.On("GetCustomer", ctx, custd.ID).Return(custd, nil)
		suite.store.On("Amounttransaction", ctx, 1200, "1234edf", "1234edfdsx").Return(nil)
		suite.store.On("GetAccountBalance", ctx, custd.ID).Twice().Return(0, nil)

		amm, err := suite.transection.Amounttransaction(ctx, 1200, "1234edf", "1234edfdsx", "1")
		assert.NoError(t, err)
		assert.Equal(t, amm, 0)

	})
	t.Run("unsuccessful transecetion insufficient Ammount", func(t *testing.T) {
		suite.SetupTest()

		suite.store.On("GetAccountBalance", ctx, custd.ID).Once().Return(1200, nil)
		suite.store.On("GetCustomer", ctx, custd.ID).Return(custd, nil)
		suite.store.On("Amounttransaction", ctx, 1300, "1234edf", "1234edfdsx").Return(nil)
		suite.store.On("GetAccountBalance", ctx, custd.ID).Twice().Return(0, nil)

		_, err := suite.transection.Amounttransaction(ctx, 1300, "1234edf", "1234edfdsx", "1")
		assert.EqualError(t, err, transaction.ErrInSufficientAmmount.Error())

	})
	t.Run("unsuccessful transecetion invalid user", func(t *testing.T) {
		suite.SetupTest()

		suite.store.On("GetAccountBalance", ctx, custd.ID).Once().Return(1200, nil)
		suite.store.On("GetCustomer", ctx, custd.ID).Return(cust2, nil)
		suite.store.On("Amounttransaction", ctx, 1200, "1234edf", "1234edfdsx").Return(nil)
		suite.store.On("GetAccountBalance", ctx, custd.ID).Twice().Return(0, nil)

		_, err := suite.transection.Amounttransaction(ctx, 1200, "1234edf", "1234edfdsx", "1")
		assert.EqualError(t, err, transaction.ErrUnAuthorize.Error())
		fmt.Printf(err.Error(), "message")
	})
}
