package domain

type AccountReference int
type LoanApplicationStatus string

type LoanApplicationState struct {
	id      int
	status  LoanApplicationStatus
	account AccountReference
}

type LoanApplication interface {
	ID() int
	Status() LoanApplicationStatus
	Account() AccountReference
	Approve()
	Reject()
}

func NewLoanApplication(accountReference AccountReference) LoanApplication {
	return &LoanApplicationState{
		id:      -1,
		status:  "pending",
		account: accountReference,
	}
}

func (las *LoanApplicationState) ID() int {
	return las.id
}

func (las *LoanApplicationState) Status() LoanApplicationStatus {
	return las.status
}

func (las *LoanApplicationState) Account() AccountReference {
	return las.account
}

func (las *LoanApplicationState) Approve() {
	las.status = "approved"
}

func (las *LoanApplicationState) Reject() {
	las.status = "rejected"
}
