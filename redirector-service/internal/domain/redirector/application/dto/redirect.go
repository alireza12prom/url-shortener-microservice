package dto

type RedirectInput struct {
	Hash string `uri:"hash" binding:"required,max=15"`
}
