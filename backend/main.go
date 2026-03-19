package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "text/plain")

		//сервер должен отвечать на "Hello from Effective Mobile!"
		if _, err := io.WriteString(w, "Hello from Effective Mobile!"); err != nil {
			log.Printf("write 'Hello from Effective Mobile!': %v", err)
		}
	})

	// this endpoint is used for health checks, it should return "OK" with status code 200
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		if _, err := io.WriteString(w, "OK"); err != nil {
			log.Printf("write 'OK': %v", err)
		}
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("server started on :8080")

	log.Fatal(server.ListenAndServe())
}
