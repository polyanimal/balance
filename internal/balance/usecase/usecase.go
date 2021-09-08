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

func (uc * BalanceUseCase) AlterFunds(id string, funds int, currency string) {

}
