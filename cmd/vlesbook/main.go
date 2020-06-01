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

	//
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := mux.NewRouter()
	r.HandleFunc("/encrypt", handlers.EncryptionHandler)

	r.HandleFunc("/decrypt", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Надо расшифровать")
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
	//
}
