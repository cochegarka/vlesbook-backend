package handlers

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"vlesbook/pkg/des"
)

func DecryptionHandler(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)

	cipherText, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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

	plainText, err := des.Decryption(cipherText, []byte(key), []byte(iv))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain;charset=UTF-8")
	_, _ = w.Write(plainText)
}
