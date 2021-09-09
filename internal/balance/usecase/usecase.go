package usecase

import (
	"Balance/internal/balance"
	"errors"
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
	var operation string

	if funds >= 0 {
		operation = "funds added"
	} else {
		operation = "funds taken"
	}

	if err := uc.repo.RecordTransaction(operation, id, "", funds, t); err != nil {
		return errors.New("error while recording a transaction")
	}

	return uc.repo.AlterFunds(id, funds)
}

func (uc * BalanceUseCase) TransferFunds(idFrom, idTo string, funds int) error {
	t := time.Now()

	//Предполагается что пользователь может только перевести свои средства другому пользователю, а не наоборот
	if funds < 0 {
		return errors.New("invalid operation")
	}

	if err := uc.repo.RecordTransaction("funds transfer", idFrom, idTo, funds, t); err != nil {
		return errors.New("error while recording a transaction")
	}

	return uc.repo.TransferFunds(idFrom, idTo, funds)
}


