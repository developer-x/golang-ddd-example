package account

import "example.com/greetings/app/account/domain"

type AccountServiceState struct {
	repo domain.AccountRepository
}

type AccountService interface {
	CreateAccount(request AccountCreateRequest) AccountView
	GetAccount(id int) AccountView
}

func NewAccountService(repo domain.AccountRepository) AccountService {
	return &AccountServiceState{repo: repo}
}

func (a *AccountServiceState) CreateAccount(request AccountCreateRequest) AccountView {
	return AccountView{}
}

func (a *AccountServiceState) GetAccount(id int) AccountView {
	//return a.repo.FindById(id)
	return AccountView{}
}
