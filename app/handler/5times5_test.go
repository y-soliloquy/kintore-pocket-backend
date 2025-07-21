package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/y-soliloquy/kintore-pocket-backend/app/handler"
)

func TestFiveTimesFiveHandler_Handle(t *testing.T) {
	// MAX重量を100kgとしてリクエストがされた前提
	reqBody := map[string]int{"weight": 100}
	body, _ := json.Marshal(reqBody)

	// テスト用にディレクトリとファイルを作成する
	jsonTemplate := `[
		{ "set": 1, "percent": 0.75, "reps": 5 },
		{ "set": 2, "percent": 0.75, "reps": 5 },
		{ "set": 3, "percent": 0.75, "reps": 5 },
		{ "set": 4, "percent": 0.75, "reps": 5 },
		{ "set": 5, "percent": 0.75, "reps": 5 }
	]`
	dir := "testdata"
	path := filepath.Join(dir, "test_5x5.json")
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatalf("failed to create testdata dir: %v", err)
	}
	defer os.RemoveAll(dir)

	if err := os.WriteFile(path, []byte(jsonTemplate), 0644); err != nil {
		t.Fatalf("failed to write test json: %v", err)
	}
	defer os.Remove(path)

	req := httptest.NewRequest(http.MethodPost, "/5times5", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler := handler.NewFiveTimesFiveHandler(path)
	handler.Handle(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("unexpected status: got %d, want %d", rr.Code, http.StatusOK)
	}

	var menus []struct {
		Set    int `json:"set"`
		Weight int `json:"weight"`
		Reps   int `json:"rep"`
	}
	if err := json.Unmarshal(rr.Body.Bytes(), &menus); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}

	// テストが固定のテンプレートを前提にしていることをコメントする
	// data/5x5.json の中身が以下の通りである前提:
	/*
		[
			{ "set": 1, "percent": 0.75, "reps": 5 },
			{ "set": 2, "percent": 0.75, "reps": 5 },
			{ "set": 3, "percent": 0.75, "reps": 5 },
			{ "set": 4, "percent": 0.75, "reps": 5 },
			{ "set": 5, "percent": 0.75, "reps": 5 }
		]
	*/
	if len(menus) != 5 {
		t.Errorf("expected 2 sets, got %d", len(menus))
	}
	if menus[0].Weight != 75 {
		t.Errorf("unexpected weight: got %d, want %d", menus[0].Weight, 75)
	}
	t.Logf("レスポンスボディ: %s", rr.Body.String())
}
