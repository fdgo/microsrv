package hmac_sha1

import (
	"crypto/hmac"
	"crypto/sha1"
)

func Sha1(message []byte) []byte {
	h := sha1.New()
	h.Write(message)
	return h.Sum(nil)
}
func HmacSha1(message, secret []byte) []byte {
	h := hmac.New(sha1.New, secret)
	h.Write(message)
	return h.Sum(nil)
}
