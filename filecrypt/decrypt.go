package filecrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"os"
)

func Decrypt(filePath string, password []byte) error {
	cipherText, err := readFile(filePath)
	if err != nil {
		return err
	}

	cipherText, nonce := extractNonceFromCipherText(cipherText)

	derivedKey := generateDerivedKey(password, nonce)

	block, err := aes.NewCipher(derivedKey)
	if err != nil {
		return err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	plainText, err := aesgcm.Open(nil, nonce, cipherText, nil)
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

func extractNonceFromCipherText(cipherText []byte) ([]byte, []byte) {
	nonce := cipherText[len(cipherText)-12:]
	cipherText = cipherText[:len(cipherText)-12]
	return cipherText, nonce
}
