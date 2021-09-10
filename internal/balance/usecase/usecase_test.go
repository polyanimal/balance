package usecase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/polyanimal/balance/internal/balance/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdvertisingUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockRepository(ctrl)
	uc := NewBalanceUC(repo)

	testId := "testId"
	balance := 1000
	testErr := errors.New("test error")

	t.Run("GetBalance", func(t *testing.T) {
		repo.EXPECT().GetBalance(testId).Return(balance, nil)
		newBalance, err := uc.GetBalance(testId, "")
		assert.Equal(t, balance, newBalance)
		assert.NoError(t, err)
	})

	t.Run("AlterFunds", func(t *testing.T) {
		repo.EXPECT().AlterFunds(testId, balance).Return(nil)
		repo.EXPECT().RecordTransaction("funds added", testId, "", balance, gomock.Any()).Return(nil)
		err := uc.AlterFunds(testId, balance)
		assert.NoError(t, err)
	})

	t.Run("AlterFunds-fail", func(t *testing.T) {
		repo.EXPECT().AlterFunds(testId, -100).Return(testErr)
		err := uc.AlterFunds(testId, -100)
		assert.Error(t, err)
	})

	t.Run("TransferFunds", func(t *testing.T) {
		repo.EXPECT().TransferFunds(testId, "someguy", balance).Return(nil)
		repo.EXPECT().RecordTransaction("funds transfer", testId, "someguy", balance, gomock.Any()).Return(nil)
		err := uc.TransferFunds(testId, "someguy", balance)
		assert.NoError(t, err)
	})
}
