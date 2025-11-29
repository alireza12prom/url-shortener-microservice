package lib

type Event struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Context       string `json:"context"`
	Payload       any    `json:"payload"`
	DateTime      string `json:"date_time"`
	CorrelationID string `json:"correlation_id"`
}

func (e *Event) AggregateID() string {
	return e.ID
}
