package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type RequestBodyDiagnosis struct {
}

type DiagnosisHandler struct {
	// dbやconfigを入れて拡張できるようにして多く
	// db *sql.DB
}

func NewDiagnosisHandler() *DiagnosisHandler {
	return &DiagnosisHandler{}
}

func (h *DiagnosisHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var req RequestBodyDiagnosis
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("failed to decode request: %v", err)
		http.Error(w, "DiagnosisHandler: Invalid input", http.StatusBadRequest)
		return
	}

}
