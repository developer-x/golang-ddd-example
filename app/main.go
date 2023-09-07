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

	accountRepo := domain.NewAccountRepository(sqliteDB)
	accountService := account.NewAccountService(accountRepo)
	accountController := account.NewController(accountService)
	api.GET("account", accountController.GetAccount)
	err := engine.Run(":8080")
	if err != nil {
		log.Fatal("Starting server failed", err)
	}
}