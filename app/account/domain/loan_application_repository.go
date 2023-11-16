package domain

import (
	"context"
	"database/sql"
	"example.com/greetings/app/utils"
	"github.com/jmoiron/sqlx"
)

type LoanApplicationRepositoryState struct {
	db *sqlx.DB
}

type DatabaseLoanApplicationState struct {
	ID      int                   `db:"id"`
	STATUS  LoanApplicationStatus `db:"status"`
	ACCOUNT AccountReference      `db:"account_fk"`
}

type LoanApplicationRepository interface {
	FindByAccountAndLoanId(
		ctx context.Context,
		account AccountReference,
		loanId int) (LoanApplication, error)
	FindAllByAccount(
		ctx context.Context,
		account AccountReference,
		pageReqest utils.PageRequest,
	) (utils.PageResponse[LoanApplication], error)
	Save(ctx context.Context, LoanApplication LoanApplication) (LoanApplication, error)
}

func NewLoanApplicationRepository(db *sqlx.DB) LoanApplicationRepository {
	return &LoanApplicationRepositoryState{db: db}
}

func (cfg *LoanApplicationRepositoryState) FindByAccountAndLoanId(
	ctx context.Context,
	account AccountReference,
	loanId int,
) (LoanApplication, error) {
	loanApplicationState := DatabaseLoanApplicationState{}
	if err := cfg.db.GetContext(
		ctx,
		&loanApplicationState,
		"SELECT id, status, account_fk FROM loan_applications WHERE account_fk=? and id=?;",
		account,
		loanId,
	); err != nil {
		return &LoanApplicationState{}, err
	}
	return &LoanApplicationState{
		id:      loanApplicationState.ID,
		status:  loanApplicationState.STATUS,
		account: loanApplicationState.ACCOUNT,
	}, nil
}

func (cfg *LoanApplicationRepositoryState) FindAllByAccount(
	ctx context.Context,
	account AccountReference,
	pageReqest utils.PageRequest,
) (utils.PageResponse[LoanApplication], error) {
	content := make([]LoanApplication, 0)
	limit := pageReqest.PageSize
	offset := pageReqest.Page * limit
	//orderBy := pageReqest.SortBy
	rows, err := cfg.db.QueryContext(
		ctx,
		"SELECT id, status, account_fk FROM loan_applications WHERE account_fk=? LIMIT ? OFFSET  ?;",
		account,
		limit,
		offset,
	)
	if err != nil {
		return utils.PageResponse[LoanApplication]{}, err
	}
	var id, account_fk int
	var status string
	for rows.Next() {
		err := rows.Scan(&id, &status, &account_fk)
		if err != nil {
			return utils.PageResponse[LoanApplication]{}, err
		}
		content = append(content, &LoanApplicationState{
			id:      id,
			status:  LoanApplicationStatus(status),
			account: AccountReference(account_fk),
		})
	}
	var count int
	if err := cfg.db.GetContext(
		ctx,
		&count,
		"SELECT count(*) FROM loan_applications WHERE account_fk=?;",
		account,
	); err != nil {
		return utils.PageResponse[LoanApplication]{}, err
	}
	return utils.NewPageResponse(
		content,
		pageReqest.Page,
		pageReqest.PageSize,
		count/pageReqest.PageSize,
	), nil
}

func (cfg *LoanApplicationRepositoryState) Save(ctx context.Context, loanApplication LoanApplication) (LoanApplication, error) {
	tx, err := cfg.db.Beginx()
	if err != nil {
		return &LoanApplicationState{}, err
	}
	var result sql.Result
	var id int
	if loanApplication.ID() < 0 {
		result, err = tx.ExecContext(
			ctx,
			`INSERT INTO loan_applications(status, account_fk) VALUES(?, ?);`,
			loanApplication.Status(),
			loanApplication.Account(),
		)
		if err != nil {
			_ = tx.Rollback()
			return &LoanApplicationState{}, err
		}
		insertId, err := result.LastInsertId()
		if err != nil {
			_ = tx.Rollback()
			return &LoanApplicationState{}, err
		}
		id = int(insertId)
	} else {
		result, err = tx.ExecContext(
			ctx,
			`UPDATE loan_applications SET status=? WHERE id=?;`,
			loanApplication.Status(),
			loanApplication.ID(),
		)
		if err != nil {
			_ = tx.Rollback()
			return &LoanApplicationState{}, err
		}
		id = loanApplication.ID()
	}
	err = tx.Commit()
	if err != nil {
		return &LoanApplicationState{}, err
	}
	return &LoanApplicationState{
		id:      id,
		status:  loanApplication.Status(),
		account: loanApplication.Account(),
	}, nil
}
