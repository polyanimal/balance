package balance

import "time"

type Repository interface {
	GetBalance(id string) (int, error)
	AlterFunds(id string, funds int) error
	TransferFunds(idFrom, idTo string, funds int) error
	RecordTransaction(operation, idFrom, idTo string, funds int, t time.Time) error
}
