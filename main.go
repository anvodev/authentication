package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})
	mux.HandleFunc("/basic-auth", basicAuthHandler)

	srv := &http.Server{
		Addr:    ":4000",
		Handler: mux,
	}
	fmt.Println("Running server on port :4000")
	err := srv.ListenAndServe()
	srv.ErrorLog.Fatal(err)
}

func basicAuthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Vary", "Authorization") // tell the cache that the response is vary base on the request Authorization header
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		fmt.Fprintf(w, "Anonymous User")
		return
	}
	encodedCredentials := strings.TrimPrefix(authorizationHeader, "Basic ")
	decodedCredentials, err := base64.StdEncoding.DecodeString(encodedCredentials)
	if err != nil {
		fmt.Fprintf(w, "Cannot decode authorization header")
		return
	}
	credentials := string(decodedCredentials)
	username, password, ok := extractUsernameAndPassword(credentials)

	if !ok {
		fmt.Fprintf(w, "Cannot extract username and password")
		return
	}

	if !authenticate(username, password) {
		fmt.Fprintf(w, "Wrong username or password")
		return
	}

	fmt.Fprintf(w, "Authentication succeed! Welcome %s!", username)
}

func extractUsernameAndPassword(credentials string) (username string, password string, ok bool) {
	parts := strings.SplitN(string(credentials), ":", 2)
	if len(parts) != 2 {
		return "", "", false
	}
	return parts[0], parts[1], true
}

func authenticate(username, password string) bool {
	return username == "alice@example.com" && password == "pa55word"
}
