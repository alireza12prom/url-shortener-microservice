package http

import (
	"github.com/gin-gonic/gin"
	exceptions "github.com/shortener-service/internal/common/exceptions"
	"github.com/shortener-service/internal/domain/shortener/application/commands"
	"github.com/shortener-service/internal/domain/shortener/application/dto"
	"github.com/shortener-service/internal/domain/shortener/domain/interfaces"
)

type ShortenerHandler struct {
	ShortenerService interfaces.ShortenerService
}

func (s *ShortenerHandler) ShortenRandom(ctx *gin.Context) {
	var input dto.ShortenRandomInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		exception := exceptions.NewBusinessException(
			exceptions.ValidationFailed,
			map[string]any{"detail": err.Error()},
		)
		ctx.Error(exception)
		return
	}

	userId, _ := ctx.Get("userId")
	command := commands.NewShortenRandomCommand(
		userId.(string),
		input.Length,
		input.Endpoint,
	)

	result, err := s.ShortenerService.ShortenRandom(command)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"result": result})
}

func (s *ShortenerHandler) ShortenCustom(ctx *gin.Context) {
	var input dto.ShortenCustomInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		exception := exceptions.NewBusinessException(
			exceptions.ValidationFailed,
			map[string]any{"detail": err.Error()},
		)
		ctx.Error(exception)
		return
	}

	userId, _ := ctx.Get("userId")
	command := commands.NewShortenCustomCommand(
		userId.(string),
		input.Hash,
		input.Endpoint,
	)

	result, err := s.ShortenerService.ShortenCustom(command)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, gin.H{"result": result})
}
