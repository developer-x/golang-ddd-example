package account

import (
	"example.com/greetings/app/account/domain"
	"example.com/greetings/app/utils"
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
	ApplyForALoan(ctx *gin.Context)
	ApproveLoan(ctx *gin.Context)
	GetLoanApplications(ctx *gin.Context)
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

type RenameCommand struct {
	Name string `json:"name"`
}

func (c *AccountControllerState) RenameAccount(ctx *gin.Context) {
	idParm := ctx.Param("id")
	id, err := strconv.Atoi(idParm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	renameCommand := RenameCommand{}
	if err := ctx.ShouldBindJSON(&renameCommand); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing request"})
		return
	}
	result, err := c.service.RenameAccount(ctx, id, renameCommand.Name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": result})
}

func (c *AccountControllerState) ApplyForALoan(ctx *gin.Context) {
	idParm := ctx.Param("id")
	id, err := strconv.Atoi(idParm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	result, err := c.service.ApplyForALoan(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": result})
}

func (c *AccountControllerState) ApproveLoan(ctx *gin.Context) {
	accountIdParm := ctx.Param("id")
	accountId, err := strconv.Atoi(accountIdParm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	loanIdParm := ctx.Param("loanId")
	loanId, err := strconv.Atoi(loanIdParm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	result, err := c.service.ApproveLoan(ctx, accountId, loanId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": result})
}

func (c *AccountControllerState) GetLoanApplications(ctx *gin.Context) {
	idParm := ctx.Param("id")
	id, err := strconv.Atoi(idParm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	pageRequest := utils.PageRequest{}
	if err := ctx.BindQuery(&pageRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	result, err := c.service.GetLoanApplicationsForAccount(ctx, id, pageRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": result})
}
