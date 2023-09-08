package domain

type AccountState struct {
	id   int
	name string
}

type Account interface {
	ID() int
	Name() string
	Rename(name string)
}

func CreateAccount(request AccountCreateRequest) Account {
	return &AccountState{id: -1, name: request.Name}
}

func (a *AccountState) ID() int {
	return a.id
}

func (a *AccountState) Name() string {
	return a.name
}

func (a *AccountState) Rename(name string) {
	a.name = name
}
