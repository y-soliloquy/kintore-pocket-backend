package handler_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/y-soliloquy/kintore-pocket-backend/app/handler"
)

func TestHelloHandler_Handle(t *testing.T) {
	h := handler.NewHelloHandler()

	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	w := httptest.NewRecorder()

	h.Handle(w, req)

	res := w.Result()
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	got := strings.TrimSpace(string(body))
	want := "hello world"
	if got != want {
		t.Errorf("expected body %q, got %q", want, got)
	}
}
