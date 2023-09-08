package account

import "example.com/greetings/app/account/domain"

type AccountView struct {
	ID   int
	Name string
}

func AccountViewFrom(account domain.Account) AccountView {
	return AccountView{ID: account.ID(), Name: account.Name()}
}
