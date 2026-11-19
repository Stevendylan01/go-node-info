package main

import (
	_ "embed"
	"io"
	"log"
	"net/http"
	"time"
)

//go:embed index.html
var indexHTML string      

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(w, indexHTML)
	})

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("Listening on %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
