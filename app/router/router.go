package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/y-soliloquy/kintore-pocket-backend/app/handler"
)

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

func NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // TODO: 本番環境（Vercelなど）ができたら追加する必要がある
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}))

	routes := []Route{
		{Method: http.MethodGet, Path: "/healthz", Handler: handler.NewHealthzHandler().Handle},
		{Method: http.MethodGet, Path: "/hello", Handler: handler.NewHelloHandler().Handle},
		{Method: http.MethodGet, Path: "/goodbye", Handler: handler.NewGoodbyeHandler().Handle},
		{Method: http.MethodPost, Path: "/training_menu", Handler: handler.NewTrainingMenuHandler("data").Handle},
		{Method: http.MethodGet, Path: "/questions", Handler: handler.NewQuestionsHandler("data").Handle},
		{Method: http.MethodPost, Path: "/diagnosis", Handler: handler.NewDiagnosisHandler().Handle},
		{Method: http.MethodGet, Path: "/reference", Handler: handler.NewReferenceHandler().Handle},
	}

	for _, route := range routes {
		r.Method(route.Method, route.Path, route.Handler)
	}

	return r
}
