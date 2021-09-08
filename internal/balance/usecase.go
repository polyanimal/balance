package balance

type UseCase interface {
	AlterFunds(id string, funds int, currency string)
}
