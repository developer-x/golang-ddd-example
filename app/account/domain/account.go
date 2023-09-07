package domain

type AccountState struct {
	Name string `json:"name"`
}

type Account interface {
	GetName() string
}

func (a *AccountState) GetName() string {
	return a.Name
}
