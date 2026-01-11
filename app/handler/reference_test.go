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
		name         string
		filename     string
		jsonTemplate string
		wantStatus   int
		wantLength   int
		wantTitles   []string
	}{
		{
			name:     "ベンチプレス解説動画",
			filename: "movies.json",
			jsonTemplate: `[
				{ "url": "", "title": "ベンチプレス解説"},
				{ "url": "", "title": "デッドリフト解説"} 
			]`,
			wantStatus: http.StatusOK,
			wantLength: 2,
			wantTitles: []string{
				"ベンチプレス解説",
				"デッドリフト解説",
			},
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

			h := handler.NewReferenceHandler(dir)

			req := httptest.NewRequest(http.MethodGet, "/reference", nil)
			rec := httptest.NewRecorder()

			h.Handle(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			if res.StatusCode != tt.wantStatus {
				t.Fatalf("status = %d, want %d", res.StatusCode, tt.wantStatus)
			}

			var got []handler.MovieInfos
			if err := json.NewDecoder(res.Body).Decode(&got); err != nil {
				t.Fatalf("failed to decode response json: %v", err)
			}

			for i, wantTitle := range tt.wantTitles {
				if got[i].Title != wantTitle {
					t.Fatalf("got[%d].Title = %q, want %q", i, got[i].Title, wantTitle)
				}
			}
		})
	}
}
