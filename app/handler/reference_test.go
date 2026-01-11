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

func TestReferenceHandler(t *testing.T) {
	tests := []struct {
		name       string
		setup      func(t *testing.T, dir string)
		wantStatus int
		wantLen    int
		wantTitles []string
	}{
		{
			name: "success: movies.jsonを返す",
			setup: func(t *testing.T, dir string) {
				t.Helper()
				jsonBody := `[
					{ "url": "https://example.com/bench", "title": "ベンチプレス解説" },
					{ "url": "https://example.com/deadlift", "title": "デッドリフト解説" }
				]`
				path := filepath.Join(dir, "movies.json")
				if err := os.WriteFile(path, []byte(jsonBody), 0644); err != nil {
					t.Fatalf("failed to write movies.json: %v", err)
				}
			},
			wantStatus: http.StatusOK,
			wantLen:    2,
			wantTitles: []string{"ベンチプレス解説", "デッドリフト解説"},
		},
		{
			name: "failure: movies.jsonが存在せず500エラー",
			setup: func(t *testing.T, dir string) {
				t.Helper()
				// 何もしない（ファイルを作らない）
			},
			wantStatus: http.StatusInternalServerError,
		},
		{
			name: "failure: movies.jsonが壊れていて500エラー",
			setup: func(t *testing.T, dir string) {
				t.Helper()
				invalid := `[{ "url": "https://example.com", "title": "ok" },` // 末尾カンマで壊す
				path := filepath.Join(dir, "movies.json")
				if err := os.WriteFile(path, []byte(invalid), 0644); err != nil {
					t.Fatalf("failed to write invalid movies.json: %v", err)
				}
			},
			wantStatus: http.StatusInternalServerError,
		},
	}

	dir := "testdata"
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatalf("failed to create testdata dir: %v", err)
	}
	defer os.RemoveAll(dir)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			tt.setup(t, dir)

			h := handler.NewReferenceHandler(dir)

			req := httptest.NewRequest(http.MethodGet, "/reference", nil)
			rec := httptest.NewRecorder()

			h.Handle(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			if res.StatusCode != tt.wantStatus {
				t.Fatalf("status = %d, want %d", res.StatusCode, tt.wantStatus)
			}

			if tt.wantStatus == http.StatusOK {
				gotCT := res.Header.Get("Content-Type")
				if gotCT != "application/json; charset=utf-8" {
					t.Fatalf("Content-Type = %q, want %q", gotCT, "application/json; charset=utf-8")
				}

				var got []handler.MovieInfos
				if err := json.NewDecoder(res.Body).Decode(&got); err != nil {
					t.Fatalf("failed to decode response json: %v", err)
				}

				if len(got) != tt.wantLen {
					t.Fatalf("len = %d, want %d", len(got), tt.wantLen)
				}

				for i, wantTitle := range tt.wantTitles {
					if got[i].Title != wantTitle {
						t.Fatalf("got[%d].Title = %q, want %q", i, got[i].Title, wantTitle)
					}
				}
			}
		})
	}
}
