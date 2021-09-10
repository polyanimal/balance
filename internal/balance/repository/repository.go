package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgconn"
	pgx "github.com/jackc/pgx/v4"
	"github.com/polyanimal/balance/internal/models"
	"time"
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
	var funds int
	sqlStatement := `
        SELECT funds FROM mdb.users 
        WHERE id=$1
    `
	err := r.db.
		QueryRow(context.Background(), sqlStatement, id).
		Scan(&funds)

	if err != nil {
		return 0, errors.New("user with specified id doesn't exist")
	}

	return funds, nil
}

func (r *BalanceRepository) createUser(id string, funds int) error {
	sqlStatement := `
		INSERT INTO mdb.users (id, funds)
		VALUES ($1, $2)
    `
	_, err := r.db.Exec(context.Background(), sqlStatement, id, funds)

	if err != nil {
		return err
	}

	return nil
}

func (r *BalanceRepository) updateUserBalance(id string, funds int) error {
	var oldBalance int

	sqlStatement := `
		SELECT funds FROM mdb.users
		WHERE id=$1
    `
	err := r.db.QueryRow(context.Background(), sqlStatement, id).Scan(&oldBalance)

	if err != nil {
		return err
	}

	if oldBalance+funds >= 0 {
		sqlStatement = `
			UPDATE mdb.users
			SET funds=$2
			WHERE id=$1
		`
		_, err = r.db.Exec(context.Background(), sqlStatement, id, oldBalance+funds)
		if err != nil {
			return err
		}
	} else {
		return errors.New("insufficient funds")
	}

	return nil
}

func (r *BalanceRepository) AlterFunds(id string, funds int) error {
	var exists bool
	sqlStatement := `
       select exists(select 1 FROM mdb.users WHERE id=$1)
    `

	r.db.
		QueryRow(context.Background(), sqlStatement, id).
		Scan(&exists)

	count := 0
	if exists {
		count = 1
	}

	switch {
	case count == 0 && funds >= 0:
		if err := r.createUser(id, funds); err != nil {
			return err
		}
		break
	case count == 0 && funds < 0:
		return errors.New("user doesn't exist")
	default:
		if err := r.updateUserBalance(id, funds); err != nil {
			return err
		}
	}

	return nil
}

func (r *BalanceRepository) TransferFunds(idFrom, idTo string, funds int) error {
	var count int
	sqlStatement := `
        SELECT COUNT(*) FROM mdb.users 
        WHERE id=$1
    `
	err := r.db.
		QueryRow(context.Background(), sqlStatement, idFrom).
		Scan(&count)

	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("user doesn't exist")
	}

	err = r.db.
		QueryRow(context.Background(), sqlStatement, idTo).
		Scan(&count)

	if err != nil {
		return err
	}

	if count == 0 {
		if err = r.createUser(idTo, 0); err != nil {
			return err
		}
	}

	var balanceFrom, balanceTo int
	sqlStatement = `
		SELECT funds FROM mdb.users
		WHERE id=$1
    `
	if err = r.db.QueryRow(context.Background(), sqlStatement, idFrom).Scan(&balanceFrom); err != nil {
		return err
	}
	if err = r.db.QueryRow(context.Background(), sqlStatement, idTo).Scan(&balanceTo); err != nil {
		return err
	}

	if balanceFrom-funds >= 0 {
		sqlStatement = `
			UPDATE mdb.users
			SET funds=$2
			WHERE id=$1
		`
		_, err = r.db.Exec(context.Background(), sqlStatement, idFrom, balanceFrom-funds)
		if err != nil {
			return err
		}

		_, err = r.db.Exec(context.Background(), sqlStatement, idTo, balanceTo+funds)
		if err != nil {
			return err
		}
	} else {
		return errors.New("insufficient funds")
	}

	return nil
}

func (r *BalanceRepository) RecordTransaction(operation, idFrom, idTo string, funds int, t time.Time) error {
	sqlStatement := `
		INSERT INTO mdb.transactions (user_id_from, user_id_to, comment, creation_date, funds)
		VALUES ($1, $2, $3, $4, $5)
    `

	_, err := r.db.Exec(context.Background(), sqlStatement, idFrom, idTo, operation, t, funds)

	if err != nil {
		return err
	}

	return nil
}

func (r *BalanceRepository) GetTransactions(userId, order, sort string, page, perPage int) ([]models.Transaction, error) {
	var sqlStatement string

	sqlStatement = `
        SELECT id, user_id_from, user_id_to, comment, creation_date, funds
        FROM mdb.transactions
        WHERE user_id_from=$1 OR user_id_to=$1 
        ORDER BY 
			CASE WHEN $2 = 'funds' AND $3 = 'ASC'  THEN funds END ASC,
			CASE WHEN $2 = 'funds' AND $3 = 'DESC' THEN funds END DESC,
			CASE WHEN $2 = 'creation_date' AND $3 = 'ASC'  THEN creation_date END ASC,
			CASE WHEN $2 = 'creation_date' AND $3 = 'DESC' THEN creation_date END DESC
		LIMIT $4 OFFSET $5 
	`

	rows, err := r.db.Query(context.Background(), sqlStatement, userId, sort, order, perPage, (page-1)*perPage)
	if err != nil {
		return nil, err
	}

	transactions := make([]models.Transaction, 0)

	for rows.Next() {
		t := models.Transaction{}
		if err := rows.Scan(&t.Id, &t.IdFrom, &t.IdTo, &t.Comment, &t.DateCreate, &t.Funds); err != nil {
			return nil, err
		}

		transactions = append(transactions, t)
	}

	return transactions, nil
}
