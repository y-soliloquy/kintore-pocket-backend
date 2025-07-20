package handler

import (
	"log"
	"net/http"
)

type HelloHandler struct {
	// dbやconfigを入れて拡張できるようにして多く
	// db *sql.DB
}

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

func (h *HelloHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("hello world"))
	if err != nil {
		log.Printf("HelloHandler: failed to write response: %v", err)
	}
}
