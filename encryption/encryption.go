package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var initializationVector = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func newAESCipherBlock() (cipher.Block, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't load from .env file", err)
	}

	mySecret, ok := os.LookupEnv("MYSECRET")
	if !ok {
		log.Fatal("No MYSECRET found")
	}

	block, err := aes.NewCipher([]byte(mySecret))

	if err != nil {
		return nil, err
	}

	return block, nil
}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// Encrypt method is to encrypt or hide any classified text
func Encrypt(text, MySecret string) (string, error) {
	block, err := newAESCipherBlock()
	if err != nil {
		return "", err
	}

	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, initializationVector)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)

	return Encode(cipherText), nil
}

func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

// Decrypt method is to extract back the encrypted text
func Decrypt(text, MySecret string) (string, error) {
	block, err := newAESCipherBlock()
	if err != nil {
		return "", err
	}

	cipherText := Decode(text)
	cfb := cipher.NewCFBDecrypter(block, initializationVector)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)

	return string(plainText), nil
}
