package main

import (
    "net/http"
    "encoding/json"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {  
    w.Header().Set("Content-type", "application/json")

    dat, err := json.Marshal(payload)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Error encoding payload"))
        return
    }

    w.WriteHeader(code)
    w.Write(dat)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
    payload := struct {
        Error string `json:"error"`
    }{
        Error: msg,
    }
    respondWithJSON(w, code, payload)
}
