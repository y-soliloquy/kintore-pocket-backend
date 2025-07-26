package handler

import (
	"log"
	"net/http"
)

type QuestionsHandler struct {
	// dbやconfigを入れて拡張できるようにして多く
	// db *sql.DB
}

func NewQuestionsHandler() *QuestionsHandler {
	return &QuestionsHandler{}
}

func (h *QuestionsHandler) Handle(w http.ResponseWriter, r *http.Response) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("questions"))
	if err != nil {
		log.Printf("HelloHandler: failed to write response: %v", err)
	}
}
