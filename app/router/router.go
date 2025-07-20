package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/y-soliloquy/kintore-pocket-backend/app/handler"
)

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

func NewRouter() http.Handler {
	r := chi.NewRouter()

	routes := []Route{
		{Method: http.MethodGet, Path: "/hello", Handler: handler.NewHelloHandler().Handle},
		{Method: http.MethodGet, Path: "/goodbye", Handler: handler.NewGoodbyeHandler().Handle},
	}

	for _, route := range routes {
		r.Method(route.Method, route.Path, route.Handler)
	}

	return r
}
