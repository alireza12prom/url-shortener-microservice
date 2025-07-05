package commands

type CaptureEventCommand struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Context       string `json:"context"`
	Payload       any    `json:"payload"`
	DateTime      string `json:"date_time"`
	CorrelationID string `json:"correlation_id"`
}

func NewCaptureEventCommand(
	id,
	name,
	context,
	correlationId string,
	payload any,
	datetime string,
) *CaptureEventCommand {
	return &CaptureEventCommand{
		ID:            id,
		Name:          name,
		Context:       context,
		Payload:       payload,
		DateTime:      datetime,
		CorrelationID: correlationId,
	}
}
