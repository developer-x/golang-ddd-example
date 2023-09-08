package account

import (
	"context"
	"example.com/greetings/app/account/domain"
)

type AccountServiceState struct {
	repo domain.AccountRepository
}

type AccountService interface {
	CreateAccount(ctx context.Context, request AccountCreateRequest) (AccountView, error)
	GetAccount(ctx context.Context, id int) (AccountView, error)
}

func NewAccountService(repo domain.AccountRepository) AccountService {
	return &AccountServiceState{repo: repo}
}

func (a *AccountServiceState) CreateAccount(ctx context.Context, request AccountCreateRequest) (AccountView, error) {
	return AccountView{}, nil
}

func (a *AccountServiceState) GetAccount(ctx context.Context, id int) (AccountView, error) {
	//return a.repo.FindById(ctx, id), nil
	return AccountView{}, nil
}
