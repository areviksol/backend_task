package main

import (
	"sync"

	"github.com/areviksol/backend_task/controller"
	"github.com/areviksol/backend_task/eventbus"
	"github.com/areviksol/backend_task/model"
	"github.com/areviksol/backend_task/view"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	model := model.NewModel("Initial data")

	eventBus := eventbus.NewEventBus()

	controller := controller.NewController(model, view.NewView(), eventBus, &wg)

	go controller.Run()

	// Simulate an event where the model is updated
	eventBus.Publish("update_model", "New data")

	wg.Wait()
}
