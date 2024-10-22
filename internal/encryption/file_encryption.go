package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"os"
)

const GGMP_HEADER = "GGMP-FILE-ENCRYPTED"

func EncryptFile(filename string, password []byte) error {
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

	fullContent := append([]byte(GGMP_HEADER), salt...)
	fullContent = append(fullContent, ciphertext...)

	err = os.Remove(filename)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, fullContent, 0644)
}

func DecryptFile(filename string, secretKey []byte) error {
	// get encrypted content
	content, err := VerifyGGMPFile(filename)
	if err != nil {
		return err
	}

	// extract salt and ciphertext
	salt, ciphertext := content[:16], content[16:]

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
	err = os.Remove(filename)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, plaintext, 0644)
	if err != nil {
		return err
	}

	err = os.Chmod(filename, 0644)
	if err != nil {
		return err
	}

	return nil
}

// To check whether the file is a valid GGMP file
// If it returns an error, it means the file is not a
// valid GGMP file
func VerifyGGMPFile(file string) (content []byte, err error) {
	content, err = os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	headerLen := len(GGMP_HEADER)
	if len(content) < headerLen || string(content[:headerLen]) != GGMP_HEADER {
		return nil, errors.New("file is not a valid GGMP file")
	}

	content = content[headerLen:]

	return content, nil

}
