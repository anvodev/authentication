package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello World")
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})
	srv := &http.Server{
		Addr:    ":4000",
		Handler: mux,
	}
	fmt.Println("Running server on port :4000")
	err := srv.ListenAndServe()
	srv.ErrorLog.Fatal(err)
}
