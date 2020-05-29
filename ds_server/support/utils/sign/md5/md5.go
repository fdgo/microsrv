package md5

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
)

func HashForPwd(salt, pwd string) string {
	m5 := md5.New()
	m5.Write([]byte(pwd))
	m5.Write([]byte(string(salt)))
	st := m5.Sum(nil)
	return hex.EncodeToString(st)
}
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func HmacMd5(message, secret []byte) []byte {
	h := hmac.New(md5.New, secret)
	h.Write(message)
	return h.Sum(nil)
}