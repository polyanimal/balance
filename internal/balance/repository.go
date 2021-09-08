package balance

type Repository interface {
	GetBalance(id string) (int, error)
	AlterFunds(id string, funds int, currency string) error
	TransferFunds(idFrom, idTo string, funds int, currency string) error
}
