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
	tests := []struct {
		name         string
		filename     string
		jsonTemplate string
		wantLength   int
		wantWeight   int
	}{
		{
			name:     "5x5",
			filename: "test_5x5.json",
			jsonTemplate: `[
				{ "set": 1, "percent": 0.75, "reps": 5 },
				{ "set": 2, "percent": 0.75, "reps": 5 },
				{ "set": 3, "percent": 0.75, "reps": 5 },
				{ "set": 4, "percent": 0.75, "reps": 5 },
				{ "set": 5, "percent": 0.75, "reps": 5 }
			]`,
			wantLength: 5,
			wantWeight: 75,
		},
		{
			name:     "piramid",
			filename: "test_piramid.json",
			jsonTemplate: `[
				{ "set": 1, "percent": 0.60, "reps": 8 },
				{ "set": 2, "percent": 0.70, "reps": 6 },
				{ "set": 3, "percent": 0.80, "reps": 4 },
				{ "set": 4, "percent": 0.90, "reps": 2 },
				{ "set": 5, "percent": 0.70, "reps": 100 }
			]`,
			wantLength: 5,
			wantWeight: 60, // 100 * 0.60 = 60
		},
	}

	dir := "testdata"
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatalf("failed to create testdata dir: %v", err)
	}
	defer os.RemoveAll(dir)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := filepath.Join(dir, tt.filename)
			if err := os.WriteFile(path, []byte(tt.jsonTemplate), 0644); err != nil {
				t.Fatalf("failed to write test json: %v", err)
			}

			reqBody := map[string]int{"weight": 100}
			body, _ := json.Marshal(reqBody)
			req := httptest.NewRequest(http.MethodPost, "/5times5?template="+tt.filename, bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			handler := handler.NewFiveTimesFiveHandler(dir)
			handler.Handle(rr, req)

			if rr.Code != http.StatusOK {
				t.Fatalf("unexpected status: got %d, want %d", rr.Code, http.StatusOK)
			}

			var menus []struct {
				Set    int `json:"set"`
				Weight int `json:"weight"`
				Reps   int `json:"reps"`
			}
			if err := json.Unmarshal(rr.Body.Bytes(), &menus); err != nil {
				t.Fatalf("invalid JSON: %v", err)
			}

			if len(menus) != tt.wantLength {
				t.Errorf("expected %d sets, got %d", tt.wantLength, len(menus))
			}
			if menus[0].Weight != tt.wantWeight {
				t.Errorf("unexpected weight: got %d, want %d", menus[0].Weight, tt.wantWeight)
			}
		})
	}
}
