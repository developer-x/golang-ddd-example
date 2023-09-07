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
