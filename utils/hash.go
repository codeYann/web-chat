// Package utils defines a bunch of utility functions like hash and other structures.
package utils

import (
	"crypto/sha256"
)

// SHA256Encoder returns a hash of a given password.
func SHA256Encoder(password string) [32]byte {
	hashPassword := sha256.Sum256([]byte(password))
	return hashPassword
}
