package handler

import (
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
	w.Write([]byte("hello world"))
}
