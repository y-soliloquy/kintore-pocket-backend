package router_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/y-soliloquy/kintore-pocket-backend/app/router"
)

func TestRouter(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
		ecpectedBody   string
	}{
		{
			name:           "healthz",
			method:         http.MethodGet,
			path:           "/healthz",
			expectedStatus: http.StatusOK,
			ecpectedBody:   "health ok",
		},
		{
			name:           "hello",
			method:         http.MethodGet,
			path:           "/hello",
			expectedStatus: http.StatusOK,
			ecpectedBody:   "hello world",
		},
		{
			name:           "goodbye",
			method:         http.MethodGet,
			path:           "/goodbye",
			expectedStatus: http.StatusOK,
			ecpectedBody:   "Goodbye",
		},
	}

	r := router.NewRouter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rep := httptest.NewRequest(tt.method, tt.path, nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, rep)

			res := w.Result()
			defer res.Body.Close()

			if res.StatusCode != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, res.StatusCode)
			}

			body, _ := io.ReadAll(res.Body)
			if strings.TrimSpace(string(body)) != tt.ecpectedBody {
				t.Errorf("expected body %q, got %q", tt.ecpectedBody, body)
			}
		})
	}
}
