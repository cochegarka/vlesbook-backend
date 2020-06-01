package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"vlesbook/pkg/des"
)

//обработчик
func EncryptionHandler(w http.ResponseWriter, r *http.Request) {
	plainText, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Bad Request with plain text")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	key := r.URL.Query().Get("key")

	iv := r.URL.Query().Get("iv")
	if len(iv) == 0 {
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
