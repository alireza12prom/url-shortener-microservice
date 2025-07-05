package dto

type CaptureEventInput struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Context  string `json:"context" binding:"required"`
	Payload  any    `json:"payload" binding:"required"`
	DateTime string `json:"dateTime" binding:"required"`
}
