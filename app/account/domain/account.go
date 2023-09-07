package domain

type AccountState struct {
	Name string `json:"name"`
}

type Account interface {
	Rename(name string)
}

func (a *AccountState) Rename(name string) {
	a.Name = name
}
