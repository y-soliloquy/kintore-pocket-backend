package handler

import (
	"encoding/json"
	"log"
	"net/http"

	diagnosis "github.com/y-soliloquy/kintore-pocket-backend/app/handler/util"
)

type Answers struct {
	Answer string
}

type RequestBodyDiagnosis struct {
	Answers []string `json:"answers"`
}

type ResponseDiagnosis struct {
	Type           string   `json:"type"`
	Recomendations []string `json:"recommendations"`
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

	result := diagnosis.Diagnose(req.Answers)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Printf("DiagnosisHandler: failed to encode response: %v", err)
	}
}
