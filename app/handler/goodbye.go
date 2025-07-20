package handler

import (
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
	w.Write([]byte("Goodbye"))
}
