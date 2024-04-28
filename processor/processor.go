package processor

import (
	"log"

	"github.com/areviksol/backend_task/eventbus"
	"github.com/areviksol/backend_task/model"
)

type Processor interface {
	Process(identifier string) error
}

type HTTPProcessor struct {
	Model    *model.Model
	EventBus *eventbus.EventBus
}

func NewHTTPProcessor(model *model.Model, eventBus *eventbus.EventBus) *HTTPProcessor {
	return &HTTPProcessor{
		Model:    model,
		EventBus: eventBus,
	}
}

func (h *HTTPProcessor) Process(identifier string) error {
	exists, err := h.Model.CheckRecord(identifier)
	if err != nil {
		return err
	}

	if !exists {
		log.Printf("Record with identifier %s does not exist", identifier)
		h.EventBus.Publish("record_not_found", identifier)
	}

	return nil
}
