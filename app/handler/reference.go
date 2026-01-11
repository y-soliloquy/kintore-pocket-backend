package handler

import (
	"log"
	"net/http"
)

// 使うかわからないが拡張できるようにしておく
type ReferenceHandler struct{}

func NewReferenceHandler() *ReferenceHandler {
	return &ReferenceHandler{}
}

func (h *ReferenceHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("reference!!"))
	if err != nil {
		log.Printf("ReferenceHandler: failed to write response: %v", err)
	}
}
