package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func handleRequests(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/ups":
		tracking := r.URL.Query().Get("tracking")

		if tracking == "" {
			http.Error(w, "Tracking code missing", http.StatusBadRequest)
			return
		}
		resp := QueryUPS(tracking)
		defer resp.Body.Close()

		for k, v := range resp.Header {
			w.Header().Set(k, v[0])
		}
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)

	default:
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}

func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func main() {
	finalHandler := http.HandlerFunc(handleRequests)
	mux := http.NewServeMux()
	mux.Handle("/", enableCors(finalHandler))

	log.Println("Server running")
	port := os.Getenv("PORT")

	if port == "" {
		port = "54321"
	}

	err := http.ListenAndServe(":"+port, mux)

	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
