package handler

import "net/http"

type DiagnosisHandler struct {
	// dbやconfigを入れて拡張できるようにして多く
	// db *sql.DB
}

func NewDiagnosisHandler() *DiagnosisHandler {
	return &DiagnosisHandler{}
}

func (h *DiagnosisHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
