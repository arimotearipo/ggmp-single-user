package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"os"
)

func EncryptFile(filename, newFilename string, key []byte) error {
	// get file content
	plaintext, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// Generate a new AES cipher using key
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// create nonce also kown as IV
	initializationVector := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(initializationVector); err != nil {
		return err
	}

	// encrypt data
	ciphertext := gcm.Seal(initializationVector, initializationVector, plaintext, nil)

	// write encrypted data to new file
	if newFilename == "" {
		newFilename = "backup_" + filename
	}

	return os.WriteFile(newFilename, ciphertext, 0644)
}

func DecryptFile(filename string, key []byte) error {
	// get encrypted content
	ciphertext, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// generate AES block using key
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// Extract the nonce size from GCM and the actual nonce from the ciphertext
	ivSize := gcm.NonceSize()
	iv, ciphertext := ciphertext[:ivSize], ciphertext[ivSize:]

	// decrypt data
	plaintext, err := gcm.Open(nil, iv, ciphertext, nil)
	if err != nil {
		return err
	}

	// rewrite file with decrypted data
	os.Remove(filename)
	return os.WriteFile(filename, plaintext, 0644)
}
