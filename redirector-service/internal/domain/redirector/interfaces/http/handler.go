package http

import (
	"github.com/gin-gonic/gin"
	"github.com/shortener-service/internal/common/exceptions"
	"github.com/shortener-service/internal/domain/redirector/application/commands"
	"github.com/shortener-service/internal/domain/redirector/application/dto"
	"github.com/shortener-service/internal/domain/redirector/interfaces"
)

type RedirectorHandler struct {
	RedirectorService interfaces.RedirectorService
}

func (Self *RedirectorHandler) Redirect(ctx *gin.Context) {
	var input dto.RedirectInput

	if err := ctx.ShouldBindUri(&input); err != nil {
		exception := exceptions.NewBusinessException(
			exceptions.ValidationFailed,
			map[string]any{"detail": err.Error()},
		)
		ctx.Error(exception)
		return
	}

	command := commands.NewRedirectCommand(input.Hash)

	result, err := Self.RedirectorService.Redirect(command)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Redirect(301, result.Endpoint)
}
