package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"log"

	"golang.org/x/crypto/pbkdf2"
)

type Encryption struct {
	InitalizationVector []byte
	Block               cipher.Block
	secret              []byte
}

func newAESCipherBlock(secret []byte) (cipher.Block, error) {
	block, err := aes.NewCipher(secret)

	if err != nil {
		return nil, err
	}

	return block, nil
}

func encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// Encrypt method is to encrypt or hide any classified text
func (c *Encryption) Encrypt(text string) (string, error) {
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
func (c *Encryption) Decrypt(text string) (string, error) {
	cipherText := decode(text)
	cfb := cipher.NewCFBDecrypter(c.Block, c.InitalizationVector)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)

	return string(plainText), nil
}

func generateSalt() []byte {
	salt := make([]byte, 16)
	rand.Read(salt)

	return salt
}

func deriveKey(password, salt []byte) []byte {
	return pbkdf2.Key(password, salt, 1000000, 32, sha256.New)
}

func NewEncryption(masterPassword []byte) *Encryption {
	salt := generateSalt()

	secret := deriveKey(masterPassword, salt)

	InitalizationVector := make([]byte, aes.BlockSize)
	rand.Read(InitalizationVector)

	Block, err := newAESCipherBlock(secret)
	if err != nil {
		log.Fatal("Fail to create blocl")
	}

	return &Encryption{
		InitalizationVector,
		Block,
		secret,
	}
}
