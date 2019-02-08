package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

type Proxy struct{}

func NewProxy() *Proxy { return &Proxy{} }

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
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

func main() {
	proxy := NewProxy()
	log.Println("Server running")

	port := os.Getenv("PORT")

	if port == "" {
		port = ":54321"
	}

	err := http.ListenAndServeTLS(":54321", "localhost.crt", "localhost.key", proxy)

	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
