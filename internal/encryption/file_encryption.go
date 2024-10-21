package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"os"
)

func EncryptFile(filename, newFilename string, password []byte) error {
	// get file content
	plaintext, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	salt, err := GenerateSalt()
	if err != nil {
		return err
	}

	key := DeriveKey(password, salt)

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

	return os.WriteFile(newFilename, append(salt, ciphertext...), 0644)
}

func DecryptFile(filename string, secretKey []byte) error {
	// get encrypted content
	saltAndCipher, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// extract salt and ciphertext
	salt, ciphertext := saltAndCipher[:16], saltAndCipher[16:]

	key := DeriveKey(secretKey, salt)

	// generate AES block using secret key
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// decrypt data
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err
	}

	// rewrite file with decrypted data
	os.Remove(filename)
	err = os.WriteFile(filename, plaintext, 0644)
	if err != nil {
		return err
	}
	return nil
}
