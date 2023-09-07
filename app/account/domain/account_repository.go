package domain

import "github.com/jmoiron/sqlx"

type AccountRepositoryState struct {
	Db *sqlx.DB
}

type AccountRepository interface {
	FindById(id int) Account
}

func NewAccountRepository(db *sqlx.DB) AccountRepository {
	return &AccountRepositoryState{Db: db}
}

func (cfg *AccountRepositoryState) FindById(id int) Account {
	return &AccountState{Name: "foobar"}
}
