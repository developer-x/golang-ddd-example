package account

import (
	"example.com/greetings/app/account/domain"
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
	RenameAccount(ctx *gin.Context)
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
	result, err := c.service.GetAccount(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": result})
}

func (c *AccountControllerState) CreateAccount(ctx *gin.Context) {
	accountCreateRequest := domain.AccountCreateRequest{}
	if err := ctx.ShouldBindJSON(&accountCreateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing request"})
		return
	}
	result, err := c.service.CreateAccount(ctx, accountCreateRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": result})
}

func (c *AccountControllerState) RenameAccount(ctx *gin.Context) {
	idParm := ctx.Param("id")
	id, err := strconv.Atoi(idParm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	body := gin.H{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing request"})
		return
	}
	name, ok := body["name"].(string)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing request"})
		return
	}
	result, err := c.service.RenameAccount(ctx, id, name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": result})
}
