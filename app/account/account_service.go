package account

import (
	"context"
	"example.com/greetings/app/account/domain"
	"example.com/greetings/app/utils"
)

type AccountServiceState struct {
	accountRepo         domain.AccountRepository
	loanApplicationRepo domain.LoanApplicationRepository
}

type AccountService interface {
	CreateAccount(ctx context.Context, request domain.AccountCreateRequest) (AccountView, error)
	GetAccount(ctx context.Context, id int) (AccountView, error)
	RenameAccount(ctx context.Context, id int, name string) (AccountView, error)
	GetLoanApplicationsForAccount(
		ctx context.Context,
		id int,
		request utils.PageRequest,
	) (utils.PageResponse[LoanApplicationView], error)
}

func NewAccountService(
	accountRepo domain.AccountRepository,
	loanApplicationRepo domain.LoanApplicationRepository,
) AccountService {
	return &AccountServiceState{
		accountRepo:         accountRepo,
		loanApplicationRepo: loanApplicationRepo,
	}
}

func (a *AccountServiceState) CreateAccount(
	ctx context.Context,
	request domain.AccountCreateRequest,
) (AccountView, error) {
	account := domain.CreateAccount(request)
	savedAccount, err := a.accountRepo.Save(ctx, account)
	if err != nil {
		return AccountView{}, err
	}
	return AccountViewFrom(savedAccount), nil
}

func (a *AccountServiceState) GetAccount(
	ctx context.Context,
	id int,
) (AccountView, error) {
	foundAccount, err := a.accountRepo.FindById(ctx, id)
	if err != nil {
		return AccountView{}, err
	}
	return AccountViewFrom(foundAccount), nil
}

func (a *AccountServiceState) RenameAccount(
	ctx context.Context,
	id int,
	name string,
) (AccountView, error) {
	foundAccount, err := a.accountRepo.FindById(ctx, id)
	if err != nil {
		return AccountView{}, err
	}
	foundAccount.Rename(name)
	savedAccount, err := a.accountRepo.Save(ctx, foundAccount)
	if err != nil {
		return AccountView{}, err
	}
	return AccountViewFrom(savedAccount), nil
}

func (a *AccountServiceState) GetLoanApplicationsForAccount(
	ctx context.Context,
	id int,
	request utils.PageRequest,
) (utils.PageResponse[LoanApplicationView], error) {
	loanApplications, err := a.loanApplicationRepo.FindAllByAccount(
		ctx,
		domain.AccountReference(id),
		request,
	)
	if err != nil {
		return utils.PageResponse[LoanApplicationView]{}, err
	}
	views := make([]LoanApplicationView, 0)
	for _, loanApplication := range loanApplications.Content {
		views = append(views, LoanApplicationViewFrom(loanApplication))
	}
	return utils.NewPageResponse(
		views,
		loanApplications.Page,
		loanApplications.PageSize,
		loanApplications.PageCount,
	), nil
}
