package main

import (
	"log"
	"net/http"

	"github.com/y-soliloquy/kintore-pocket-backend/app/config"
	"github.com/y-soliloquy/kintore-pocket-backend/app/db"
	"github.com/y-soliloquy/kintore-pocket-backend/app/router"
)

func main() {
	cfg := config.LoadDBConfig()

	dbClient, dbErr := db.NewDBClient(cfg.DatabaseURL)
	if dbErr != nil {
		log.Fatalf("db init failed: %v", dbErr)
	}
	defer dbClient.Close()

	log.Println("application started")

	r := router.NewRouter()

	log.Println("Listening on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
