package server

import (
	"net/http"

	"github.com/areviksol/backend_task/controller"
)

type Server struct {
	Controller *controller.Controller
}

func NewServer(controller *controller.Controller) *Server {
	return &Server{Controller: controller}
}

func (s *Server) Run() error {
	http.HandleFunc("/", s.Controller.HandleRequest)
	return http.ListenAndServe(":5000", nil)
}
