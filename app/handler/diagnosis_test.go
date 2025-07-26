package handler_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/y-soliloquy/kintore-pocket-backend/app/handler"
)

func TestDiagnosisHandler_Handle_Success(t *testing.T) {
	h := handler.NewDiagnosisHandler()

	reqBody := handler.RequestBodyDiagnosis{
		Answers: []string{"A", "B", "A", "B"}, // AとBが同率で最大
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/diagnosis", bytes.NewReader(body))
	w := httptest.NewRecorder()

	h.Handle(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 OK, got %d", res.StatusCode)
	}

	if ct := res.Header.Get("Content-Type"); ct != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", ct)
	}

	// レスポンス確認
	var decoded map[string]interface{}
	resBody, _ := io.ReadAll(res.Body)
	if err := json.Unmarshal(resBody, &decoded); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if _, ok := decoded["results"]; !ok {
		t.Errorf("expected key 'results' in response")
	}
}

func TestDiagnosisHandler_Handle_InvalidJSON(t *testing.T) {
	h := handler.NewDiagnosisHandler()

	invalidJSON := []byte(`{ "answers": [1, 2, 3] }`) // 数値は不正（string配列を期待）

	req := httptest.NewRequest(http.MethodPost, "/diagnosis", bytes.NewReader(invalidJSON))
	w := httptest.NewRecorder()

	h.Handle(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected 400 Bad Request, got %d", res.StatusCode)
	}
}
