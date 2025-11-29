package http

import (
	"github.com/gin-gonic/gin"
	exceptions "github.com/shortener-service/internal/common/exceptions"
	"github.com/shortener-service/internal/domain/account/application/commands"
	"github.com/shortener-service/internal/domain/account/application/dto"
	"github.com/shortener-service/internal/domain/account/application/queries"
	"github.com/shortener-service/internal/domain/account/interfaces"
)

type AccountHandler struct {
	AccountService interfaces.AccountService
}

func (h *AccountHandler) RegisterAccount(ctx *gin.Context) {
	var input dto.RegisterAccountInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		exception := exceptions.NewBusinessException(
			exceptions.ValidationFailed,
			map[string]any{"detail": err.Error()},
		)
		ctx.Error(exception)
		return
	}

	command := commands.NewRegisterAccountCommand(&input)
	result, err := h.AccountService.RegisterAccount(&command)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"result": result})
}

func (h *AccountHandler) LoginAccount(ctx *gin.Context) {
	var input dto.LoginAccountInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		exception := exceptions.NewBusinessException(
			exceptions.ValidationFailed,
			map[string]any{"detail": err.Error()},
		)
		ctx.Error(exception)
		return
	}

	command := commands.NewLoginAccountCommand(&input)
	result, err := h.AccountService.LoginAccount(&command)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, gin.H{"result": result})
}

func (h *AccountHandler) ChangePassword(ctx *gin.Context) {
	var input dto.ChangePasswordInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		exception := exceptions.NewBusinessException(
			exceptions.ValidationFailed,
			map[string]any{"detail": err.Error()},
		)
		ctx.Error(exception)
		return
	}

	userId, _ := ctx.Get("userId")
	command := commands.NewChangePasswordCommand(
		userId.(string),
		input.NewPassword,
		input.OldPassword,
	)

	result, err := h.AccountService.ChangePassword(&command)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, gin.H{"result": result})
}

func (h *AccountHandler) ChangeUsername(ctx *gin.Context) {
	var input dto.ChangeUsernameInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		exception := exceptions.NewBusinessException(
			exceptions.ValidationFailed,
			map[string]any{"detail": err.Error()},
		)
		ctx.Error(exception)
		return
	}

	userId, _ := ctx.Get("userId")
	command := commands.NewChangeUsernameCommand(
		userId.(string),
		input.Username,
	)

	result, err := h.AccountService.ChangeUsername(&command)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, gin.H{"result": result})
}

func (h *AccountHandler) GetAccount(ctx *gin.Context) {
	var input dto.GetAccountInput
	userId, _ := ctx.Get("userId")
	input.UserID = userId.(string)

	query := queries.NewGetAccountQuery(&input)
	result, err := h.AccountService.GetAccount(&query)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, gin.H{"result": result})
}
