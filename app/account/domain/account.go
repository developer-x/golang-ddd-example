package domain

type AccountState struct {
	name string
}

type Account interface {
	Rename(name string)
}

func NewAccount() Account {
	return nil
}

func (a *AccountState) Rename(name string) {
	a.name = name
}
