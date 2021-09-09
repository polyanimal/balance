package balance

type UseCase interface {
	AlterFunds(id string, funds int) error
	GetBalance(id string, currency string) (int, error)
	TransferFunds(idFrom, idTo string, funds int) error
}
