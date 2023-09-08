package domain

import "github.com/jmoiron/sqlx"

type AccountRepositoryState struct {
	db *sqlx.DB
}

type DatabaseAccountState struct {
	ID string `db:"id"`
}

type AccountRepository interface {
	FindById(id int) Account
	Save(account Account) Account
}

func NewAccountRepository(db *sqlx.DB) AccountRepository {
	return &AccountRepositoryState{db: db}
}

func (cfg *AccountRepositoryState) FindById(id int) Account {
	return &AccountState{name: "foobar"}
}

func (cfg *AccountRepositoryState) Save(account Account) Account {
	return nil
}
