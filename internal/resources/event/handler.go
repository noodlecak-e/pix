package event

type EventHandler struct{}

func NewHandler() *EventHandler {
	return &EventHandler{}
}
