package filecrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

func Encrypt(filePath string, password []byte) error {
	plainText, err := readFile(filePath)
	if err != nil {
		return err
	}

	nonce := make([]byte, 12)

	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return err
	}

	derivedKey := generateDerivedKey(password, nonce)

	block, err := aes.NewCipher(derivedKey)
	if err != nil {
		return err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	cipherText := aesgcm.Seal(nil, nonce, plainText, nil)
	cipherText = append(cipherText, nonce...)

	dstFile, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer dstFile.Close()

	_, err = dstFile.Write(cipherText)
	if err != nil {
		return err
	}

	return nil
}
