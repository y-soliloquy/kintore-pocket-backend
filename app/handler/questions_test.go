package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/y-soliloquy/kintore-pocket-backend/app/handler"
)

func setupTestQuestionsJSON(t *testing.T) string {
	t.Helper()
	dir := "testdata"
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatalf("failed to create testdata dir: %v", err)
	}

	path := filepath.Join(dir, "questions.json")
	questions := []handler.Question{
		{
			ID:    "1",
			Title: "どんなトレーニングが好き？",
			Options: []handler.Option{
				{Label: "パンプ感が好き", Type: "A"},
				{Label: "高重量で力を出すのが好き", Type: "B"},
				{Label: "長く走るのが好き", Type: "C"},
			},
		},
	}
	data, err := json.Marshal(questions)
	if err != nil {
		t.Fatalf("failed to marshal questions: %v", err)
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		t.Fatalf("failed to write questions.json: %v", err)
	}

	return path
}

func teardownTestQuestionsJSON(path string) {
	_ = os.RemoveAll(filepath.Dir(path))
}

func TestQuestionsHandler_Handle(t *testing.T) {
	path := setupTestQuestionsJSON(t)
	defer teardownTestQuestionsJSON(path)

	req := httptest.NewRequest(http.MethodGet, "/questions", nil)
	rr := httptest.NewRecorder()

	h := handler.NewQuestionsHandler("testdata")
	h.Handle(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("unexpected status: got %d, want %d", rr.Code, http.StatusOK)
	}

	var got []handler.Question
	if err := json.Unmarshal(rr.Body.Bytes(), &got); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}

	if len(got) != 1 {
		t.Errorf("expected 1 question, got %d", len(got))
	}
}
