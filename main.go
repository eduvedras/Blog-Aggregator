package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/eduvedras/Blog-Aggregator/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main(){
	err := godotenv.Load()
	if err != nil{
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	dbURL := os.Getenv("CONN_STRING")
	
	db, err := sql.Open("postgres", dbURL)
	if err != nil{
		log.Fatal(err)
	}

	cfg := apiConfig{
		DB: database.New(db),
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/healthz", handlerReadiness)
	mux.HandleFunc("GET /v1/err", handlerErr)
	mux.HandleFunc("POST /v1/users", cfg.handlerCreateUser)
	mux.HandleFunc("GET /v1/users", cfg.handlerGetUser)

	server := &http.Server{
		Addr: ":" + port,
		Handler: mux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}