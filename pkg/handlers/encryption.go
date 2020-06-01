package handlers

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"vlesbook/pkg/des"
)

//обработчик
func EncryptionHandler(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)

	plainText, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Bad Request with plain text")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	key, ok := args["key"]
	if !ok {
		log.Println("Bad Request with keys")
		log.Println(args)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	iv, ok := args["iv"]
	if !ok {
		iv = key
	}

	cipherText, err := des.Encryption(plainText, []byte(key), []byte(iv))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain;charset=UTF-8")
	_, _ = w.Write(cipherText)
}
