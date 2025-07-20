package main

import (
	"log"
	"net/http"

	"github.com/y-soliloquy/kintore-pocket-backend/app/router"
)

func main() {
	r := router.NewRouter()

	log.Println("Listening on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
