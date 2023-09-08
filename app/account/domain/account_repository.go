package domain

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryState struct {
	db *sqlx.DB
}

type DatabaseAccountState struct {
	ID   int    `db:"id"`
	NAME string `db:"name"`
}

type AccountRepository interface {
	FindById(ctx context.Context, id int) (Account, error)
	Save(ctx context.Context, account Account) (Account, error)
}

func NewAccountRepository(db *sqlx.DB) AccountRepository {
	return &AccountRepositoryState{db: db}
}

func (cfg *AccountRepositoryState) FindById(ctx context.Context, id int) (Account, error) {
	accountState := DatabaseAccountState{}
	if err := cfg.db.GetContext(ctx, &accountState, "SELECT * FROM accounts WHERE id=?;", id); err != nil {
		return &AccountState{}, err
	}
	return &AccountState{id: accountState.ID, name: accountState.NAME}, nil
}

func (cfg *AccountRepositoryState) Save(ctx context.Context, account Account) (Account, error) {
	tx, err := cfg.db.Beginx()
	if err != nil {
		return &AccountState{}, err
	}
	var result sql.Result
	var id int
	if account.ID() < 0 {
		result, err = tx.ExecContext(
			ctx,
			`INSERT INTO accounts(name) VALUES(?);`,
			account.Name(),
		)
		if err != nil {
			_ = tx.Rollback()
			return &AccountState{}, err
		}
		insertId, err := result.LastInsertId()
		if err != nil {
			_ = tx.Rollback()
			return &AccountState{}, err
		}
		id = int(insertId)
	} else {
		result, err = tx.ExecContext(
			ctx,
			`UPDATE accounts SET name=? WHERE id=?;`,
			account.Name(),
			account.ID(),
		)
		if err != nil {
			_ = tx.Rollback()
			return &AccountState{}, err
		}
		id = account.ID()
	}
	err = tx.Commit()
	if err != nil {
		return &AccountState{}, err
	}
	return &AccountState{id: id, name: account.Name()}, nil
}
