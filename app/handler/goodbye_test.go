package handler_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/y-soliloquy/kintore-pocket-backend/app/handler"
)

func TestGoodbyeHandler_Handle(t *testing.T) {
	h := handler.NewGoodbyeHandler()

	req := httptest.NewRequest(http.MethodGet, "/goodbye", nil)
	w := httptest.NewRecorder()

	h.Handle(w, req)

	res := w.Result()
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	got := strings.TrimSpace(string(body))
	want := "Goodbye"
	if got != want {
		t.Errorf("expected body %q, got %q", want, got)
	}
}
