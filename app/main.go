package main

import (
	"example.com/greetings/app/account"
	"example.com/greetings/app/account/domain"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	engine := gin.Default()
	api := engine.Group("/api")
	sqliteDB := sqlx.MustOpen("sqlite3", ":memory:")
	sqliteDB.MustExec(`CREATE TABLE accounts (id INTEGER PRIMARY KEY, name TEXT);`)
	sqliteDB.MustExec(`CREATE TABLE loan_applications (id INTEGER PRIMARY KEY, status TEXT, account_fk INTEGER);`)

	accountRepo := domain.NewAccountRepository(sqliteDB)
	loanApplicationRepo := domain.NewLoanApplicationRepository(sqliteDB)
	accountService := account.NewAccountService(accountRepo, loanApplicationRepo)
	accountController := account.NewController(accountService)
	api.POST("accounts", accountController.CreateAccount)
	api.GET("accounts/:id", accountController.GetAccount)
	api.GET("accounts/:id/loan-applications", accountController.GetLoanApplications)
	api.POST("accounts/:id/loan-applications", accountController.ApplyForALoan)
	api.PATCH("accounts/:id/loan-applications/:loanId/approve", accountController.ApproveLoan)
	api.PATCH("accounts/:id/name", accountController.RenameAccount)
	err := engine.Run(":8080")
	if err != nil {
		log.Fatal("Starting server failed", err)
	}
}
