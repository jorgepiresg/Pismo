// Code generated by MockGen. DO NOT EDIT.
// Source: transactions.go

// Package mocksApp is a generated GoMock package.
package mocksApp

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	modelTransactions "github.com/jorgepiresg/ChallangePismo/model/transactions"
)

// MockITransactions is a mock of ITransactions interface.
type MockITransactions struct {
	ctrl     *gomock.Controller
	recorder *MockITransactionsMockRecorder
}

// MockITransactionsMockRecorder is the mock recorder for MockITransactions.
type MockITransactionsMockRecorder struct {
	mock *MockITransactions
}

// NewMockITransactions creates a new mock instance.
func NewMockITransactions(ctrl *gomock.Controller) *MockITransactions {
	mock := &MockITransactions{ctrl: ctrl}
	mock.recorder = &MockITransactionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITransactions) EXPECT() *MockITransactionsMockRecorder {
	return m.recorder
}

// Make mocks base method.
func (m *MockITransactions) Make(ctx context.Context, data modelTransactions.MakeTransaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Make", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Make indicates an expected call of Make.
func (mr *MockITransactionsMockRecorder) Make(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Make", reflect.TypeOf((*MockITransactions)(nil).Make), ctx, data)
}