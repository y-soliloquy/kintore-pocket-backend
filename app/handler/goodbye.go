package handler

import (
	"log"
	"net/http"
)

type GoodbyeHandler struct {
	// dbやconfigを入れて拡張できるようにして多く
	// db *sql.DB
}

func NewGoodbyeHandler() *GoodbyeHandler {
	return &GoodbyeHandler{}
}

func (h *GoodbyeHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Goodbye"))
	if err != nil {
		log.Printf("GoodbyeHandler: failed to write response: %v", err)
	}
}
