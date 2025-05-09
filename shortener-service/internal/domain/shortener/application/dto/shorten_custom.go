package dto

type ShortenCustomInput struct {
	Hash     string `json:"hash" binding:"required,max=15"`
	Endpoint string `json:"endpoint" binding:"required,max=200"`
}
