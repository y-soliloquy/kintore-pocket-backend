package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type MovieInfos struct {
	URL   string `json:"url"`
	TiTle string `json:"title"`
}

// 使うかわからないが拡張できるようにしておく
type ReferenceHandler struct {
	baseDir string
}

func NewReferenceHandler(path string) *ReferenceHandler {
	return &ReferenceHandler{baseDir: path}
}

func (h *ReferenceHandler) Handle(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(h.baseDir, "movies.json")

	b, err := os.ReadFile(path)
	if err != nil {
		log.Printf("ReferenceHandler: failed to read file: %v", err)
		http.Error(w, "failed to read reference json", http.StatusInternalServerError)
		return
	}

	var infos []MovieInfos
	if err := json.Unmarshal(b, &infos); err != nil {
		log.Printf("ReferenceHandler: failed to parse json: %v", err)
		http.Error(w, "invalid reference json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(infos); err != nil {
		log.Printf("ReferenceHandler: failed to write response: %v", err)
	}
}
