package dto

type ShortenRandomInput struct {
	Length   int    `json:"length" binding:"required,gte=5,lte=15"`
	Endpoint string `json:"endpoint" binding:"required,max=200"`
}
