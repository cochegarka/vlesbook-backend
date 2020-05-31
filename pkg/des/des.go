package des

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"errors"
	"strconv"
)

//процедура дешифровки
//вход: ключ, вектор инициализации. зашифрованный текст
//выход: открытый текст, сообщение об ошибке
func Decryption(key, iv, cipherText []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	sourceData := make([]byte, len(cipherText))
	blockMode.CryptBlocks(sourceData, cipherText)
	plainText := PKCS5UnPadding(sourceData)
	return plainText, nil
}

//процедура шифрования
//вход: ключ, вектор инициализации. открытый текст
//выход: шифрованный текст, ошибка
func Encryption(key, iv, plainText []byte) ([]byte, error) {
	block, err := des.NewCipher(key) //создать новый шифр
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if len(iv) != blockSize {
		return nil, errors.New("Invalid iv length (must be " + strconv.Itoa(blockSize) + ")")
	}
	alignedData := PKCS5Padding(plainText, blockSize) //выровнять данные
	blockMode := cipher.NewCBCEncrypter(block, iv)
	encrypted := make([]byte, len(alignedData))
	blockMode.CryptBlocks(encrypted, alignedData)
	return encrypted, nil
}

//процедура дозаполнения текста для краткного кодирования блоками
func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

//процедура снятия дозаполнения
func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
