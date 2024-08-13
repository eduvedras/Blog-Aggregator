package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

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

	dbQueries := database.New(db)
	cfg := apiConfig{
		DB: dbQueries,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/healthz", handlerReadiness)
	mux.HandleFunc("GET /v1/err", handlerErr)

	mux.HandleFunc("POST /v1/users", cfg.handlerCreateUser)
	mux.HandleFunc("GET /v1/users", cfg.middlewareAuth(cfg.handlerGetUser))

	mux.HandleFunc("POST /v1/feeds", cfg.middlewareAuth(cfg.handlerCreateFeed))
	mux.HandleFunc("GET /v1/feeds", cfg.handlerGetFeeds)
	
	mux.HandleFunc("POST /v1/feed_follows", cfg.middlewareAuth(cfg.handlerCreateFeedFollow))
	mux.HandleFunc("DELETE /v1/feed_follows/{feedFollowID}", cfg.middlewareAuth(cfg.handlerDeleteFeedFollow))
	mux.HandleFunc("GET /v1/feed_follows", cfg.middlewareAuth(cfg.handlerGetFeedFollowsOfUser))

	mux.HandleFunc("GET /v1/posts", cfg.middlewareAuth(cfg.handlerGetPostsByUser))

	server := &http.Server{
		Addr: ":" + port,
		Handler: mux,
	}

	const collectionConcurrency = 10
	const collectionInterval = time.Minute
	go startScraping(dbQueries, collectionConcurrency, collectionInterval)

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}