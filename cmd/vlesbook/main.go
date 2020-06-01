package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"vlesbook/pkg/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := mux.NewRouter()
	r.HandleFunc("/encrypt", handlers.EncryptionHandler).Methods("POST")
	r.HandleFunc("/decrypt", handlers.DecryptionHandler).Methods("POST")

	h := corsMiddleware(r)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), h))
	//
}

// Allows CORS for given handler.
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		next.ServeHTTP(w, r)
	})
}
