package account

import (
	"context"
	"example.com/greetings/app/account/domain"
)

type AccountServiceState struct {
	repo domain.AccountRepository
}

type AccountService interface {
	CreateAccount(ctx context.Context, request domain.AccountCreateRequest) (AccountView, error)
	GetAccount(ctx context.Context, id int) (AccountView, error)
}

func NewAccountService(repo domain.AccountRepository) AccountService {
	return &AccountServiceState{repo: repo}
}

func (a *AccountServiceState) CreateAccount(ctx context.Context, request domain.AccountCreateRequest) (AccountView, error) {
	account := domain.CreateAccount(request)
	savedAccount, err := a.repo.Save(ctx, account)
	if err != nil {
		return AccountView{}, err
	}
	return AccountViewFrom(savedAccount), nil
}

func (a *AccountServiceState) GetAccount(ctx context.Context, id int) (AccountView, error) {
	foundAccount, err := a.repo.FindById(ctx, id)
	if err != nil {
		return AccountView{}, err
	}
	return AccountViewFrom(foundAccount), nil
}
