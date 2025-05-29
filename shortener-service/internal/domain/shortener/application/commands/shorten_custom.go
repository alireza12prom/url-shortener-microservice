package commands

type ShortenCustomCommand struct {
	UserID   string
	Hash     string
	Endpoint string
}

func NewShortenCustomCommand(userId, hash, endpoint string) *ShortenCustomCommand {
	return &ShortenCustomCommand{
		UserID:   userId,
		Hash:     hash,
		Endpoint: endpoint,
	}
}

type ShortenCustomCommandOutput struct {
	ShortenURL string `json:"shorten_url"`
}
