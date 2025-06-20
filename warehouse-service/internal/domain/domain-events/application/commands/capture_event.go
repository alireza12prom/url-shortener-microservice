package commands

type CaptureEventCommand struct {
	ID            string
	Name          string
	Context       string
	Payload       any
	DateTime      string
	CorrelationID string
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
