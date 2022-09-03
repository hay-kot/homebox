package hasher

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
)

type Token struct {
	Raw  string
	Hash []byte
}

func GenerateToken() Token {
	randomBytes := make([]byte, 16)
	_, _ = rand.Read(randomBytes)

	plainText := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
	hash := HashToken(plainText)

	return Token{
		Raw:  plainText,
		Hash: hash,
	}
}

func HashToken(plainTextToken string) []byte {
	hash := sha256.Sum256([]byte(plainTextToken))
	return hash[:]
}
