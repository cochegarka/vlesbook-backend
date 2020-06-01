package handlers

import (
	"io/ioutil"
	"net/http"
	"vlesbook/pkg/des"
	"vlesbook/pkg/utils"
)

func DecryptionHandler(w http.ResponseWriter, r *http.Request) {
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

	cipherTextRepresentation := utils.ParseByteSlice(string(cipherText))

	plainText, err := des.Decryption(cipherTextRepresentation, []byte(key), []byte(iv))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain;charset=UTF-8")
	_, _ = w.Write(plainText)
}
