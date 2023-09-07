package account

import "example.com/greetings/app/account/domain"

type AccountServiceState struct {
	Repo domain.AccountRepository
}

type AccountService interface {
	GetAccount(id int) domain.Account
}

func NewAccountService(repo domain.AccountRepository) AccountService {
	return &AccountServiceState{Repo: repo}
}

func (a *AccountServiceState) GetAccount(id int) domain.Account {
	return a.Repo.FindById(id)
}
