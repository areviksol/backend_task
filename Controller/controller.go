package controller

import (
	"net/http"

	"github.com/areviksol/backend_task/eventbus"
	"github.com/areviksol/backend_task/model"
	"github.com/areviksol/backend_task/processor"
)

type Controller struct {
	Processor processor.Processor
	EventBus  *eventbus.EventBus
	Model     *model.Model
}

func NewController(processor processor.Processor, eventBus *eventbus.EventBus, model *model.Model) *Controller {
	return &Controller{
		Processor: processor,
		EventBus:  eventBus,
		Model:     model,
	}
}

func (c *Controller) HandleRequest(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	identifier := queryValues.Get("id")

	if identifier == "" {
		http.Error(w, "Missing identifier parameter", http.StatusBadRequest)
		return
	}

	exists, err := c.Model.CheckRecord(identifier)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if !exists {
		c.EventBus.Publish("record_not_found", identifier)
	}

	w.Write([]byte("OK"))
}
