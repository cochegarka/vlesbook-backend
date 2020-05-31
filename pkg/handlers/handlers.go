package handlers

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"vlesbook/pkg/des"
)

//обработчик
func EncryptionHandler(w http.ResponseWriter, r *http.Request) {
	// Настройки CORS. Важно!
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

	args := mux.Vars(r) //получить аргументы запроса

	plainText, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //если текст не считан
		return
	}

	key, ok := args["key"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	iv, ok := args["iv"]
	if !ok {
		iv = key
	}

	cipherText, err := des.Encryption(plainText, []byte(key), []byte(iv))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) //если ошибка при кодировании
		return
	}
	w.Header().Set("Content-Type", "text/plain;charset=UTF-8")
	_, _ = w.Write(cipherText)
}
