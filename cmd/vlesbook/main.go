package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := mux.NewRouter()
	r.HandleFunc("/encrypt", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Надо зашифровать")
	})

	r.HandleFunc("/decrypt", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Надо расшифровать")
	})

	// TODO:
	// Погугли "golang dotenv", разберись как работает, вынеси порт в env загружай через переменные среды (os.Env)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
