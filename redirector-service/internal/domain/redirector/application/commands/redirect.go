package commands

type RedirectCommand struct {
	Hash string
}

func NewRedirectCommand(hash string) *RedirectCommand {
	return &RedirectCommand{
		Hash: hash,
	}
}

type RedirectCommandOutput struct {
	Endpoint string `json:"endpoint"`
}
