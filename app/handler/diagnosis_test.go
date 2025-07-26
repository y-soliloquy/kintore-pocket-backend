package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/y-soliloquy/kintore-pocket-backend/app/handler"
)

func TestDiagnosisHandler_Handle(t *testing.T) {
	tests := []struct {
		name          string
		request       handler.RequestBodyDiagnosis
		expectedType  string
		expectedMenus []string
	}{
		{
			name: "type A wins",
			request: handler.RequestBodyDiagnosis{
				Answers: []string{"A", "A", "B", "C"},
			},
			expectedType:  "A",
			expectedMenus: []string{"ピラミッド法", "アセンディング法", "ディセンディング法"},
		},
		{
			name: "type C wins",
			request: handler.RequestBodyDiagnosis{
				Answers: []string{"C", "C", "A", "B"},
			},
			expectedType:  "C",
			expectedMenus: []string{"有酸素運動"},
		},
		{
			name: "type B wins",
			request: handler.RequestBodyDiagnosis{
				Answers: []string{"B", "B", "B", "C"},
			},
			expectedType:  "B",
			expectedMenus: []string{"5x5法", "3x3法"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.request)
			if err != nil {
				t.Fatalf("failed to marshal request: %v", err)
			}

			req := httptest.NewRequest(http.MethodPost, "/diagnosis", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			h := handler.NewDiagnosisHandler()
			h.Handle(rr, req)

			if rr.Code != http.StatusOK {
				t.Fatalf("unexpected status code: got %d, want %d", rr.Code, http.StatusOK)
			}

			var got handler.ResponseDiagnosis
			if err := json.Unmarshal(rr.Body.Bytes(), &got); err != nil {
				t.Fatalf("failed to unmarshal response: %v", err)
			}

			if got.Type != tt.expectedType {
				t.Errorf("unexpected type: got %s, want %s", got.Type, tt.expectedType)
			}

			if len(got.Recomendations) != len(tt.expectedMenus) {
				t.Fatalf("unexpected recommendations length: got %d, want %d", len(got.Recomendations), len(tt.expectedMenus))
			}
			for i := range got.Recomendations {
				if got.Recomendations[i] != tt.expectedMenus[i] {
					t.Errorf("recommendation[%d]: got %s, want %s", i, got.Recomendations[i], tt.expectedMenus[i])
				}
			}
		})
	}
}
