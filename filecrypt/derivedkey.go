package filecrypt

import (
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"
)

func generateDerivedKey(password []byte, nonce []byte) []byte {
	iterations := 4096
	keyLength := 32
	derivedKey := pbkdf2.Key(password, nonce, iterations, keyLength, sha256.New)
	return derivedKey
}
