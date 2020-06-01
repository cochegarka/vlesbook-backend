package handlers

import (
	"io/ioutil"
	"net/http"
	"vlesbook/pkg/des"
)

func DecryptionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

	cipherText, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	key := r.URL.Query().Get("key")
	if len(key) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	iv := r.URL.Query().Get("iv")
	if len(iv) == 0 {
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
