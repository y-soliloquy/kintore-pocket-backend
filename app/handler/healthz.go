package handler

import (
	"log"
	"net/http"
)

type HealthzHandler struct{}

func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{}
}

func (h *HealthzHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("health ok"))
	if err != nil {
		log.Printf("HealthzHandler: failed to write response: %v", err)
	}
}
