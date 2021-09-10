package usecase

import (
	"errors"
	"github.com/polyanimal/balance/internal/balance"
	"github.com/polyanimal/balance/internal/models"
	"time"
)

type BalanceUseCase struct {
	repo balance.Repository
}

func NewBalanceUC(repo balance.Repository) *BalanceUseCase {
	return &BalanceUseCase{
		repo: repo,
	}
}

func (uc * BalanceUseCase) GetBalance(id string, currency string) (int, error) {
	return uc.repo.GetBalance(id)
}

func (uc * BalanceUseCase) AlterFunds(id string, funds int) error {
	t := time.Now()

	if err := uc.repo.AlterFunds(id, funds); err != nil {
		return err
	}

	var operation string
	if funds >= 0 {
		operation = "funds added"
	} else {
		operation = "funds taken"
	}

	if err := uc.repo.RecordTransaction(operation, id, "", funds, t); err != nil {
		return errors.New("error while recording a transaction")
	}

	return nil
}

func (uc * BalanceUseCase) TransferFunds(idFrom, idTo string, funds int) error {
	t := time.Now()

	//Предполагается что пользователь может только перевести свои средства другому пользователю, а не наоборот
	if funds < 0 {
		return errors.New("invalid operation")
	}

	if err:=  uc.repo.TransferFunds(idFrom, idTo, funds); err != nil {
		return err
	}

	if err := uc.repo.RecordTransaction("funds transfer", idFrom, idTo, funds, t); err != nil {
		return errors.New("error while recording a transaction")
	}

	return nil
}

func (uc * BalanceUseCase) GetTransactions(request models.TransactionsRequest) ([]models.Transaction, error) {
	return uc.repo.GetTransactions(request.UserId, request.Order, request.Sort, request.Page, request.PerPage)
}

