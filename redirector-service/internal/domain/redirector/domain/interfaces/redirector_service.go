package interfaces

import "github.com/shortener-service/internal/domain/redirector/application/commands"

type RedirectorService interface {
	Redirect(command *commands.RedirectCommand) (*commands.RedirectCommandOutput, error)
}
