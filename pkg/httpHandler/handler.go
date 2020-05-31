package httpHandler

import (
	"github.com/gorilla/mux"
	"net/http"
	"vlesbook/pkg/des"
)

//обработчик
func EncryptionHandler(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)              //получить аргументы запроса
	var plainText []byte             //буфер открытого текста
	_, err := r.Body.Read(plainText) //считать тело запроса
	if err != nil {
		w.WriteHeader(400) //если текст не считан
		return
	}
	key, ok := args["key"]
	if !ok {
		w.WriteHeader(456)
		return
	}
	iv, ok := args["iv"]
	if !ok {
		w.WriteHeader(456)
		return
	}
	cipherText, err := des.Encryption(plainText, []byte(key), []byte(iv))
	if err != nil {
		w.WriteHeader(500) //если ошибка при кодировании
		return
	}
	w.Header().Set("Content-Type", "text/plain;charset=UTF-8")
	w.Write(cipherText)
}
