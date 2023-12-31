package accounts

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	mocksStore "github.com/jorgepiresg/ChallangePismo/mocks/store"
	modelAccounts "github.com/jorgepiresg/ChallangePismo/model/accounts"
	"github.com/jorgepiresg/ChallangePismo/store"
	"github.com/sirupsen/logrus"
)

func TestCreate(t *testing.T) {

	type fields struct {
		accounts *mocksStore.MockIAccounts
	}

	tests := map[string]struct {
		input    modelAccounts.Create
		expected modelAccounts.Account
		err      error
		prepare  func(f *fields)
	}{
		"should be able to create a new account": {
			input: modelAccounts.Create{
				DocumentNumber: "111.111.111-11",
			},
			prepare: func(f *fields) {
				f.accounts.EXPECT().GetByDocument(gomock.Any(), "11111111111").Times(1).Return(modelAccounts.Account{}, fmt.Errorf("any"))
				f.accounts.EXPECT().Create(gomock.Any(), modelAccounts.Create{DocumentNumber: "11111111111"}).Times(1).Return(modelAccounts.Account{
					ID:             "id",
					DocumentNumber: "11111111111",
				}, nil)
			},
			expected: modelAccounts.Account{
				ID:             "id",
				DocumentNumber: "11111111111",
			},
		},
		"should not be able to create a new account with error document number invalid caracters": {
			input: modelAccounts.Create{
				DocumentNumber: "111.111.111-AB",
			},
			prepare: func(f *fields) {},
			err:     fmt.Errorf("document number invalid"),
		},
		"should not be able to create a new account with error document number invalid": {
			input: modelAccounts.Create{
				DocumentNumber: "111",
			},
			prepare: func(f *fields) {},
			err:     fmt.Errorf("document number invalid"),
		},
		"should not be able to create a new account with error account alredy exist": {
			input: modelAccounts.Create{
				DocumentNumber: "11111111111",
			},
			prepare: func(f *fields) {
				f.accounts.EXPECT().GetByDocument(gomock.Any(), "11111111111").Times(1).Return(modelAccounts.Account{
					ID:             "id",
					DocumentNumber: "11111111111",
				}, nil)
			},
			err: fmt.Errorf("account alredy exist"),
		},
	}

	for key, tt := range tests {
		t.Run(key, func(t *testing.T) {

			ctrl := gomock.NewController(t)

			accountsMock := mocksStore.NewMockIAccounts(ctrl)

			tt.prepare(&fields{
				accounts: accountsMock,
			})

			a := New(Options{
				Store: store.Store{
					Accounts: accountsMock,
				},
				Log: logrus.New(),
			})

			res, err := a.Create(context.Background(), tt.input)

			if err != nil && err.Error() != tt.err.Error() {
				t.Errorf(`Expected err: "%s" got "%s"`, tt.err, err)
			}
			if res != tt.expected {
				t.Errorf("Expected result %v got %v", tt.expected, res)
			}

		})
	}
}

func TestGetByAccountID(t *testing.T) {
	type fields struct {
		accounts *mocksStore.MockIAccounts
	}

	tests := map[string]struct {
		input    string
		expected modelAccounts.Account
		err      error
		prepare  func(s *fields)
	}{
		"should be able to get account by id": {
			input: "id",
			prepare: func(f *fields) {
				f.accounts.EXPECT().GetByID(gomock.Any(), "id").Times(1).Return(modelAccounts.Account{ID: "id", DocumentNumber: "11111111111"}, nil)
			},
			expected: modelAccounts.Account{ID: "id", DocumentNumber: "11111111111"},
		},
		"should not be able to get account by id with error account not found": {
			input: "id",
			prepare: func(f *fields) {
				f.accounts.EXPECT().GetByID(gomock.Any(), "id").Times(1).Return(modelAccounts.Account{}, fmt.Errorf("any"))
			},
			err: fmt.Errorf("account not found"),
		},
	}

	for key, tt := range tests {
		t.Run(key, func(t *testing.T) {

			ctrl := gomock.NewController(t)

			accountsMock := mocksStore.NewMockIAccounts(ctrl)

			tt.prepare(&fields{
				accounts: accountsMock,
			})

			a := New(Options{
				Store: store.Store{
					Accounts: accountsMock,
				},
			})

			res, err := a.GetByAccountID(context.Background(), tt.input)

			if err != nil && err.Error() != tt.err.Error() {
				t.Errorf(`Expected err: "%s" got "%s"`, tt.err, err)
			}
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("Expected result %v got %v", tt.expected, res)
			}
		})
	}
}
