package balance

import "github.com/polyanimal/balance/internal/models"

type UseCase interface {
	AlterFunds(id string, funds int) error
	GetBalance(id string, currency string) (int, error)
	TransferFunds(idFrom, idTo string, funds int) error
	GetTransactions(request models.TransactionsRequest) ([]models.Transaction, error)
}
