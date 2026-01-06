package handler

import "net/http"

// 使うかわからないが拡張できるようにしておく
type ReferenceHandler struct{}

func (h *ReferenceHandler) Handle(w http.ResponseWriter, r http.Request) {
	w.WriteHeader(http.StatusOK)
}
