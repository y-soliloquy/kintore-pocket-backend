package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Option struct {
	Label string `json:"label"`
	Type  string `json:"type"` // "A"、"B"、もしくは"C"
}

type Question struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Options []Option `json:"options"`
}

type QuestionsHandler struct {
	// dbやconfigを入れて拡張できるようにして多く
	// db *sql.DB
}

func NewQuestionsHandler() *QuestionsHandler {
	return &QuestionsHandler{}
}

func (h *QuestionsHandler) Handle(w http.ResponseWriter, r *http.Response) {
	w.WriteHeader(http.StatusOK)
	questions, err := LoadQuestions("../data/questions.json")
	if err != nil {
		http.Error(w, "failed to load questions", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}

func LoadQuestions(path string) ([]Question, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("QuestionsHandler: failed to load json: %v", err)
		return nil, err
	}

	var questions []Question
	if err := json.Unmarshal(data, &questions); err != nil {
		log.Printf("QuestionsHandler: failed to unmarshal json: %v", err)
		return nil, err
	}

	return questions, nil
}
