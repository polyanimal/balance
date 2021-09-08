package usecase

import "Balance/internal/balance"

type BalanceUseCase struct {
	repo balance.Repository
}

func NewBalanceUC(repo balance.Repository) *BalanceUseCase {
	return &BalanceUseCase{
		repo: repo,
	}
}

func (uc * BalanceUseCase) GetBalance(id string) (int, error) {
	return uc.repo.GetBalance(id)
}

func (uc * BalanceUseCase) AlterFunds(id string, funds int, currency string) error {
	return uc.repo.AlterFunds(id, funds, currency)
}

func (uc * BalanceUseCase) TransferFunds(idFrom, idTo string, funds int, currency string) error {
	return uc.repo.TransferFunds(idFrom, idTo, funds, currency)
}

