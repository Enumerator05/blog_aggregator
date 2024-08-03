package main

import (
    "net/http"
    "blot_aggregator/internal/database"
    "blot_aggregator/internal/auth"
)


type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        apiKey, err := auth.GetApiKey(r.Header)
        if err != nil {
            respondWithError(w, http.StatusBadRequest, "api key not found")
            return
        }

        user, err := cfg.DB.GetUserByApiKey(r.Context(), apiKey)
        if err != nil {
            respondWithError(w, http.StatusNotFound, "could not find user")
            return
        }

        handler(w, r, user)
    }
}
