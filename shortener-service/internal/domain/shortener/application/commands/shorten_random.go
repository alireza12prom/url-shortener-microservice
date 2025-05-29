package commands

type ShortenRandomCommand struct {
	UserID   string
	Length   int
	Endpoint string
}

func NewShortenRandomCommand(userId string, length int, endpoint string) *ShortenRandomCommand {
	return &ShortenRandomCommand{
		UserID:   userId,
		Length:   length,
		Endpoint: endpoint,
	}
}

type ShortenRandomCommandOutput struct {
	ShortenURL string `json:"shorten_url"`
}
