package domain

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryState struct {
	db *sqlx.DB
}

type DatabaseAccountState struct {
	ID string `db:"id"`
}

type AccountRepository interface {
	FindById(ctx context.Context, id int) (Account, error)
	Save(ctx context.Context, account Account) (Account, error)
}

func NewAccountRepository(db *sqlx.DB) AccountRepository {
	return &AccountRepositoryState{db: db}
}

func (cfg *AccountRepositoryState) FindById(ctx context.Context, id int) (Account, error) {
	return &AccountState{name: "name"}, nil
}

func (cfg *AccountRepositoryState) Save(ctx context.Context, account Account) (Account, error) {
	return &AccountState{name: "name"}, nil
}
