package des

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"errors"
	"strconv"
)

// Decryption дешифрует данные
// Входные данные: ключ, вектор инициализации. зашифрованный текст
// Выходные данные: открытый текст, сообщение об ошибке
func Decryption(cipherData, key, iv []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	sourceData := make([]byte, len(cipherData))
	blockMode.CryptBlocks(sourceData, cipherData)
	plainText := pkcs5Unpadding(sourceData)
	return plainText, nil
}

// Encryption расшифровывает данные
// Вход: ключ, вектор инициализации. открытый текст
// Выход: шифрованный текст, ошибка
func Encryption(plainData, key, iv []byte) ([]byte, error) {
	block, err := des.NewCipher(key) //создать новый шифр
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if len(iv) != blockSize {
		return nil, errors.New("Invalid iv length (must be " + strconv.Itoa(blockSize) + ")")
	}
	alignedData := pkcs5Padding(plainData, blockSize) //выровнять данные
	blockMode := cipher.NewCBCEncrypter(block, iv)
	encrypted := make([]byte, len(alignedData))
	blockMode.CryptBlocks(encrypted, alignedData)
	return encrypted, nil
}

// Процедура дозаполнения текста для кратного кодирования блоками
func pkcs5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

// Процедура снятия дозаполнения
func pkcs5Unpadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
