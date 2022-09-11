package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

func GenerateHashString(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	hashedPassword := hex.EncodeToString(hash.Sum(nil))
	return hashedPassword
}
