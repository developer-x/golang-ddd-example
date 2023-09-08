package account

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AccountControllerState struct {
	service AccountService
}

type AccountController interface {
	GetAccount(ctx *gin.Context)
	CreateAccount(ctx *gin.Context)
}

func NewController(service AccountService) AccountController {
	return &AccountControllerState{
		service: service,
	}
}

func (c *AccountControllerState) GetAccount(ctx *gin.Context) {
	idParm := ctx.Param("id")
	id, err := strconv.Atoi(idParm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	result := c.service.GetAccount(id)
	ctx.JSON(http.StatusOK, gin.H{"response": result})
}

func (c *AccountControllerState) CreateAccount(ctx *gin.Context) {
	accountCreateRequest := AccountCreateRequest{}
	if err := ctx.ShouldBindJSON(&accountCreateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing request"})
		return
	}
	result := c.service.CreateAccount(accountCreateRequest)
	ctx.JSON(http.StatusOK, gin.H{"response": result})
}
