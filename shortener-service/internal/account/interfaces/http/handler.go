package http

import (
	"github.com/gin-gonic/gin"
	"github.com/shortener-service/internal/account/application/commands"
	"github.com/shortener-service/internal/account/application/queries"
)

type Handler struct {
	RegisterAccountCmd *commands.RegisterAccountCommand
	ChangePasswordCmd  *commands.ChangePasswordCommand
	ChangeUsernameCmd  *commands.ChangeUsernameCommand
	GetAccountQr       *queries.GetAccountQuery
}

func (h *Handler) RegisterAccount(ctx *gin.Context) {
}

func (h *Handler) ChangePassword(ctx *gin.Context) {
}

func (h *Handler) ChangeUsername(ctx *gin.Context) {
}

func (h *Handler) GetAccount(ctx *gin.Context) {
}
