package repository

import (
	"context"
	"github.com/jackc/pgconn"
	pgx "github.com/jackc/pgx/v4"
)

// PgxPoolIface Интерфейс для драйвера БД
type PgxPoolIface interface {
	Begin(context.Context) (pgx.Tx, error)
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	Ping(context.Context) error
}

type BalanceRepository struct {
	db PgxPoolIface
}

func NewBalanceRepository(database PgxPoolIface) *BalanceRepository {
	return &BalanceRepository{
		db: database,
	}
}

func (r *BalanceRepository) GetBalance(id string) (int, error) {
	return 0, nil
}

func (r *BalanceRepository) AlterFunds(id string, funds int, currency string) error {
	return nil
}

func (r *BalanceRepository) TransferFunds(idFrom, idTo string, funds int, currency string) error {
	return nil
}
