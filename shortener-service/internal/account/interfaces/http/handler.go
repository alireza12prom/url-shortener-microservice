package http

import (
	"github.com/alireza12prom/url-shortener-microservice/shortener-service/internal/account/application/commands"
	"github.com/alireza12prom/url-shortener-microservice/shortener-service/internal/account/application/queries"
	"github.com/gin-gonic/gin"
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
