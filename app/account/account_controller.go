package account

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountControllerState struct {
	service AccountService
}

type AccountController interface {
	GetAccount(ctx *gin.Context)
}

func NewController(service AccountService) AccountController {
	return &AccountControllerState{
		service: service,
	}
}

func (c *AccountControllerState) GetAccount(ctx *gin.Context) {
	result := c.service.GetAccount()
	ctx.JSON(http.StatusOK, gin.H{"response": result})
}
