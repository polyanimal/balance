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

type BalanceRepo struct {
	db PgxPoolIface
}

func NewBalanceRepository(database PgxPoolIface) *BalanceRepo {
	return &BalanceRepo{
		db: database,
	}
}


