package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})
	mux.HandleFunc("/basic-auth", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Authorization") // tell the cache that the response is vary base on the request Authorization header
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			json.NewEncoder(w).Encode(map[string]string{"user": "anonymous user"})
			return
		}
		username, password, ok := r.BasicAuth()
		if !ok {
			json.NewEncoder(w).Encode(map[string]string{"error": "auth failed"})
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"username": username, "password": password})
	})
	srv := &http.Server{
		Addr:    ":4000",
		Handler: mux,
	}
	fmt.Println("Running server on port :4000")
	err := srv.ListenAndServe()
	srv.ErrorLog.Fatal(err)
}
