// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	db "github.com/ayushjnv1/Gobank/db"
	mock "github.com/stretchr/testify/mock"
)

// Storer is an autogenerated mock type for the Storer type
type Storer struct {
	mock.Mock
}

// AmmountDeposit provides a mock function with given fields: ctx, cid, amount
func (_m *Storer) AmmountDeposit(ctx context.Context, cid string, amount int) error {
	ret := _m.Called(ctx, cid, amount)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int) error); ok {
		r0 = rf(ctx, cid, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AmmountWithdraw provides a mock function with given fields: ctx, cid, amount
func (_m *Storer) AmmountWithdraw(ctx context.Context, cid string, amount int) error {
	ret := _m.Called(ctx, cid, amount)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int) error); ok {
		r0 = rf(ctx, cid, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Amounttransection provides a mock function with given fields: ctx, amount, creditAcc, debitAcc
func (_m *Storer) Amounttransection(ctx context.Context, amount int, creditAcc string, debitAcc string) error {
	ret := _m.Called(ctx, amount, creditAcc, debitAcc)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string, string) error); ok {
		r0 = rf(ctx, amount, creditAcc, debitAcc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateCustomer provides a mock function with given fields: ctx, uid
func (_m *Storer) CreateCustomer(ctx context.Context, uid string) (db.Customer, error) {
	ret := _m.Called(ctx, uid)

	var r0 db.Customer
	if rf, ok := ret.Get(0).(func(context.Context, string) db.Customer); ok {
		r0 = rf(ctx, uid)
	} else {
		r0 = ret.Get(0).(db.Customer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *Storer) CreateUser(ctx context.Context, user db.User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, db.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCustomer provides a mock function with given fields: ctx, cid
func (_m *Storer) DeleteCustomer(ctx context.Context, cid string) (db.Customer, error) {
	ret := _m.Called(ctx, cid)

	var r0 db.Customer
	if rf, ok := ret.Get(0).(func(context.Context, string) db.Customer); ok {
		r0 = rf(ctx, cid)
	} else {
		r0 = ret.Get(0).(db.Customer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, cid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUser provides a mock function with given fields: ctx, id
func (_m *Storer) DeleteUser(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByEmail provides a mock function with given fields: ctx, email
func (_m *Storer) FindByEmail(ctx context.Context, email string) (db.User, error) {
	ret := _m.Called(ctx, email)

	var r0 db.User
	if rf, ok := ret.Get(0).(func(context.Context, string) db.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(db.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: ctx, id
func (_m *Storer) FindById(ctx context.Context, id string) (db.User, error) {
	ret := _m.Called(ctx, id)

	var r0 db.User
	if rf, ok := ret.Get(0).(func(context.Context, string) db.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCustomer provides a mock function with given fields: ctx, cid
func (_m *Storer) GetCustomer(ctx context.Context, cid string) (db.Customer, error) {
	ret := _m.Called(ctx, cid)

	var r0 db.Customer
	if rf, ok := ret.Get(0).(func(context.Context, string) db.Customer); ok {
		r0 = rf(ctx, cid)
	} else {
		r0 = ret.Get(0).(db.Customer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, cid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetammountAcc provides a mock function with given fields: ctx, id
func (_m *Storer) GetammountAcc(ctx context.Context, id string) (int, error) {
	ret := _m.Called(ctx, id)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, string) int); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListOfUser provides a mock function with given fields: ctx
func (_m *Storer) ListOfUser(ctx context.Context) ([]db.User, error) {
	ret := _m.Called(ctx)

	var r0 []db.User
	if rf, ok := ret.Get(0).(func(context.Context) []db.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePassword provides a mock function with given fields: ctx, pass, Id
func (_m *Storer) UpdatePassword(ctx context.Context, pass string, Id string) error {
	ret := _m.Called(ctx, pass, Id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, pass, Id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUser provides a mock function with given fields: ctx, user, id
func (_m *Storer) UpdateUser(ctx context.Context, user db.User, id string) error {
	ret := _m.Called(ctx, user, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, db.User, string) error); ok {
		r0 = rf(ctx, user, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewStorer interface {
	mock.TestingT
	Cleanup(func())
}

// NewStorer creates a new instance of Storer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStorer(t mockConstructorTestingTNewStorer) *Storer {
	mock := &Storer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}