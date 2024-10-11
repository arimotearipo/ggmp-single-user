package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

func deriveKey(password []byte, salt []byte) []byte {
	time := uint32(1)
	memory := uint32(64 * 1024)
	threads := uint8(4)
	keyLen := uint32(32)

	return argon2.IDKey(password, salt, time, memory, threads, keyLen)
}

func Decrypt(encrypted string, masterPassword string) (string, error) {
	// Decode the encrypted string to bytes
	decoded, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	// Extract salt, IV and encryption bytes
	salt := decoded[:16]
	initializationVector := decoded[16:28]
	encryptionBytes := decoded[28:]

	// Derive key using master password
	key := deriveKey([]byte(masterPassword), []byte(salt))

	// Create cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Create GCM
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Decrypt
	plainText, err := gcm.Open(nil, initializationVector, encryptionBytes, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}

func Encrypt(password, masterPassword string) (ciphertText string, err error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	// Derive key from master password
	key := deriveKey([]byte(masterPassword), salt)

	// Create cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Create GCM
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Create IV (also known as nonce)
	initializationVector := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(initializationVector); err != nil {
		return "", err
	}

	// Encrypt. The initializationVector will prepended to the ciphertext
	ciphertext := gcm.Seal(initializationVector, initializationVector, []byte(password), nil)

	// Combine salt and ciphertext, then encode
	return base64.StdEncoding.EncodeToString(append(salt, ciphertext...)), nil
}
