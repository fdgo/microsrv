package hmac_sha512

import (
	"crypto/hmac"
	"crypto/sha512"
)

func HmacSha512(message, secret []byte) []byte {
	h := hmac.New(sha512.New, secret)
	h.Write(message)
	return h.Sum(nil)
}
