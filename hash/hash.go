package hash

import (
	"crypto/sha512"
	"encoding/base64"
)

// HashService provides an interface for string hashing
type HashService interface {
	Calculate(str string) string
}

// HashSha512 implements SHA512 hashing
type HashSha512 struct{}

// Calculate returns a SHA512 sum for a given string
func (h *HashSha512) Calculate(str string) string {
	sha := sha512.Sum512([]byte(str))
	return base64.StdEncoding.EncodeToString(sha[:])
}
