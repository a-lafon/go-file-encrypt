package filecrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
	"os"
)

func Decrypt(filePath string, password []byte) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	cipherText, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	salt := cipherText[len(cipherText)-12:]
	cipherText = cipherText[:len(cipherText)-12]

	derivedKey := generateDerivedKey(password, salt)

	block, err := aes.NewCipher(derivedKey)
	if err != nil {
		return err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	plainText, err := aesgcm.Open(nil, salt, cipherText, nil)
	if err != nil {
		return err
	}

	dstFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = dstFile.Write(plainText)
	if err != nil {
		return err
	}

	return nil
}
