package balance

type UseCase interface {
	AlterFunds(id string, funds int, currency string) error
	GetBalance(id string) (int, error)
	TransferFunds(idFrom, idTo string, funds int, currency string) error
}
