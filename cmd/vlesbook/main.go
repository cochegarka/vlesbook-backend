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

	// TODO:
	// Погугли "golang dotenv", разберись как работает, вынеси порт в env загружай через переменные среды (os.Env)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
	//
}

/*
пример работы с dex

originalText := "Соси писюн и не психуй"
	fmt.Println(originalText)
	mytext := []byte(originalText)

	key := []byte{0xBC, 0xBC, 0xBC, 0xBC, 0xBC, 0xBC, 0xBC, 0xBC }
	iv := []byte{0xBA, 0xDC, 0xBC, 0xBC, 0xBC, 0xBC, 0xBC, 0xBC }
	cryptoText,_ := des.Encryption(key, iv, mytext)
	fmt.Println(string(cryptoText))
	plainText, _ := des.Decryption(key, iv, cryptoText)
	fmt.Println(string(plainText))
*/
