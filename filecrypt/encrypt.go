package filecrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

func Encrypt(filePath string, password []byte) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	plainText, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	salt := make([]byte, 12)

	_, err = io.ReadFull(rand.Reader, salt)
	if err != nil {
		return err
	}

	derivedKey := generateDerivedKey(password, salt)

	block, err := aes.NewCipher(derivedKey)
	if err != nil {
		return err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	cipherText := aesgcm.Seal(nil, salt, plainText, nil)
	cipherText = append(cipherText, salt...)

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
