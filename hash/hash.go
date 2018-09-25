package hash

import (
	"crypto/sha512"
	"encoding/base64"
)

type HashService interface {
	Calculate(str string) (string, error)
}

type Hash struct{}

func (h *Hash) Calculate(str string) string {
	sha := sha512.Sum512([]byte(str))
	return base64.StdEncoding.EncodeToString(sha[:])
}
