package auth

import (
    "strings"
    "errors"
    "net/http"
)

var(
    ErrMalformedHeader = errors.New("Malformed Header")
    ErrInvalidApiKey = errors.New("Invalid Api Key")
)
func GetApiKey(headers http.Header) (string, error) {
    authentiation := headers.Get("Authorization")
    if authentiation == "" {
        return "", ErrMalformedHeader
    }

    tokens := strings.Split(authentiation, " ")

    if len(tokens) != 2 {
        return "", ErrMalformedHeader
    }

    if tokens[0] != "ApiKey" {
        return "", ErrMalformedHeader
    }

    if len(tokens[1]) != 64 {
        return "", ErrInvalidApiKey
    }
    return tokens[1] ,nil
}
