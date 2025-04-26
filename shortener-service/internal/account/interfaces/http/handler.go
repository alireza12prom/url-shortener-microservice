package http

import (
	"github.com/gin-gonic/gin"
	"github.com/shortener-service/internal/account/application/commands"
	"github.com/shortener-service/internal/account/application/dto"
	"github.com/shortener-service/internal/account/application/queries"
	"github.com/shortener-service/internal/account/domain/interfaces"
)

type AccountHandler struct {
	AccountService interfaces.AccountService
}

func (h *AccountHandler) RegisterAccount(ctx *gin.Context) {
	var input dto.RegisterAccountInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	command := commands.NewRegisterAccountCommand(&input)
	result, err := h.AccountService.RegisterAccount(&command)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to register account"})
		return
	}

	ctx.JSON(200, gin.H{"result": result})
}

func (h *AccountHandler) ChangePassword(ctx *gin.Context) {
	var input dto.ChangePasswordInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	command := commands.NewChangePasswordCommand(&input)
	result, err := h.AccountService.ChangePassword(&command)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to change password"})
		return
	}

	ctx.JSON(200, gin.H{"result": result})
}

func (h *AccountHandler) ChangeUsername(ctx *gin.Context) {
	var input dto.ChangeUsernameInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	command := commands.NewChangeUsernameCommand(&input)
	result, err := h.AccountService.ChangeUsername(&command)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to change username"})
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
		ctx.JSON(500, gin.H{"error": "Failed to get account"})
		return
	}

	ctx.JSON(200, gin.H{"result": result})
}
