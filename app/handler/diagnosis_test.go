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

func TestDiagnosisHandler_Handle(t *testing.T) {
	h := handler.NewDiagnosisHandler()

	tests := []struct {
		name          string
		answers       []string
		expectedTypes []string
	}{
		{
			name:          "A only",
			answers:       []string{"A", "A", "B"},
			expectedTypes: []string{"A"},
		},
		{
			name:          "A and B tie",
			answers:       []string{"A", "B", "A", "B"},
			expectedTypes: []string{"A", "B"},
		},
		{
			name:          "B and C tie",
			answers:       []string{"B", "C", "C", "B"},
			expectedTypes: []string{"B", "C"},
		},
		{
			name:          "A and C tie",
			answers:       []string{"A", "C", "C", "A"},
			expectedTypes: []string{"A", "C"},
		},
		{
			name:          "A, B, C tie",
			answers:       []string{"A", "B", "C"},
			expectedTypes: []string{"A", "B", "C"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody := handler.RequestBodyDiagnosis{
				Answers: tt.answers,
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

			var response struct {
				Results []struct {
					Type string `json:"type"`
				} `json:"results"`
			}
			resBody, _ := io.ReadAll(res.Body)
			if err := json.Unmarshal(resBody, &response); err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}

			// 実際に返ってきたタイプ一覧を取り出し
			var gotTypes []string
			for _, r := range response.Results {
				gotTypes = append(gotTypes, r.Type)
			}

			// mapで比較（順不同対応）
			wantMap := make(map[string]bool)
			for _, t := range tt.expectedTypes {
				wantMap[t] = true
			}
			gotMap := make(map[string]bool)
			for _, t := range gotTypes {
				gotMap[t] = true
			}

			if len(wantMap) != len(gotMap) {
				t.Errorf("expected types %v, got %v", tt.expectedTypes, gotTypes)
			}
			for typ := range wantMap {
				if !gotMap[typ] {
					t.Errorf("expected type %s not found in response", typ)
				}
			}
		})
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
