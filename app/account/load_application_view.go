package account

import "example.com/greetings/app/account/domain"

type LoanApplicationView struct {
	ID     int
	STATUS string
}

func LoanApplicationViewFrom(loanApplication domain.LoanApplication) LoanApplicationView {
	return LoanApplicationView{ID: loanApplication.ID(), STATUS: string(loanApplication.Status())}
}
