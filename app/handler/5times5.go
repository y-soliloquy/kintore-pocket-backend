package handler

import (
	"log"
	"net/http"
)

type FiveTimesFiveHandler struct{}

func NewFiveTimesFiveHandler() *FiveTimesFiveHandler {
	return &FiveTimesFiveHandler{}
}

// 一旦仮で書く
func (h *FiveTimesFiveHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("ok"))
	if err != nil {
		log.Printf("FiveTimesFiveHandler: failed to write response: %v", err)
	}
}
