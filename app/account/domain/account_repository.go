package domain

import "github.com/jmoiron/sqlx"

type AccountRepositoryConfig struct {
	Db *sqlx.DB
}

type AccountRepository interface {
	FindById(id uint) Account
}

func NewAccountRepository(db *sqlx.DB) AccountRepository {
	return &AccountRepositoryConfig{Db: db}
}

func (cfg *AccountRepositoryConfig) FindById(id uint) Account {
	return &AccountState{Name: "foobar"}
}
