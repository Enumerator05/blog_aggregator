package main

import (
    "net/http"
    "github.com/joho/godotenv"
    "os"
    "log"
    _ "github.com/lib/pq"
    "blot_aggregator/internal/database"
    "database/sql"
    "fmt"
)

type apiConfig struct {
    DB *database.Queries
}

func main() {
    godotenv.Load(".env")
    port := os.Getenv("PORT")
    dbURL := os.Getenv("DB_URL")
    
    db, err := sql.Open("postgres", dbURL)

    if err != nil {
        log.Fatal(err)
    }

    dbQueries := database.New(db)

    cfg := &apiConfig{
        DB: dbQueries,
    }
    // create the server
    mux := http.NewServeMux()
    
    mux.HandleFunc("POST /v1/users", cfg.HandlerCreateUser)
    mux.HandleFunc("GET /v1/healthz", HandlerReadiness)
    mux.HandleFunc("GET /v1/err", HandlerError)
    mux.HandleFunc("GET /v1/users", cfg.HandlerGetUser)
    mux.HandleFunc("POST /v1/feeds", cfg.MiddlewareAuth(cfg.HandlerCreateFeed))
    mux.HandleFunc("GET /v1/feeds", cfg.HandlerGetFeeds)
    mux.HandleFunc("DELETE /v1/feed_follows/{feed_follow_id}", cfg.MiddlewareAuth(cfg.HandlerDeleteFeedFollow))
    mux.HandleFunc("POST /v1/feed_follows/", cfg.MiddlewareAuth(cfg.HandlerCreateFeedFollow))
    mux.HandleFunc("GET /v1/feed_follows/", cfg.MiddlewareAuth(cfg.HandlerGetFeedFollows))

    srv := http.Server{
        Addr:   ":" + port,
        Handler:    mux,
    }

    log.Print(fmt.Sprintf("starting on port %s", port))
    log.Fatal(srv.ListenAndServe())
}
