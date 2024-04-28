package controller

import (
	"fmt"
	"sync"

	"github.com/areviksol/backend_task/eventbus"
	"github.com/areviksol/backend_task/model"
	"github.com/areviksol/backend_task/view"
)

// Controller struct manages interactions between Model and View
type Controller struct {
	model    *model.Model
	view     *view.View
	eventBus *eventbus.EventBus
	wg       *sync.WaitGroup
}

// NewController creates a new instance of Controller
func NewController(model *model.Model, view *view.View, eventBus *eventbus.EventBus, wg *sync.WaitGroup) *Controller {
	return &Controller{
		model:    model,
		view:     view,
		eventBus: eventBus,
		wg:       wg,
	}
}

// Run starts the controller
func (c *Controller) Run() {
	defer c.wg.Done()

	// Listen for events
	for {
		select {
		case data := <-c.eventBus.Subscribe("update_model"):
			c.model.SetData(data.(string))
			fmt.Println("Model updated:", c.model.GetData())
		}
	}
}
