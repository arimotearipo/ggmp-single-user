package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Cryption struct {
	InitalizationVector []byte
	Block               cipher.Block
	secret              string
}

func newAESCipherBlock(secret string) (cipher.Block, error) {
	block, err := aes.NewCipher([]byte(secret))

	if err != nil {
		return nil, err
	}

	return block, nil
}

func encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// Encrypt method is to encrypt or hide any classified text
func (c *Cryption) Encrypt(text string) (string, error) {
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(c.Block, c.InitalizationVector)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)

	return encode(cipherText), nil
}

func decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

// Decrypt method is to extract back the encrypted text
func (c *Cryption) Decrypt(text string) (string, error) {
	cipherText := decode(text)
	cfb := cipher.NewCFBDecrypter(c.Block, c.InitalizationVector)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)

	return string(plainText), nil
}

func NewCrypt() *Cryption {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't load from .env file", err)
	}

	secret, ok := os.LookupEnv("MYSECRET")
	if !ok {
		log.Fatal("No MYSECRET found")
	}

	InitalizationVector := make([]byte, aes.BlockSize)
	rand.Read(InitalizationVector)

	Block, err := newAESCipherBlock(secret)
	if err != nil {
		log.Fatal("Fail to create blocl")
	}

	return &Cryption{
		InitalizationVector,
		Block,
		secret,
	}
}
