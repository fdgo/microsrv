package base64

import (
	"encoding/base64"
	"net/url"
)

func UrlEncode(source string) string {
	return url.QueryEscape(source)
}
func UrlDecode(source string) (string, error) {
	return url.QueryUnescape(source)
}
func Base64(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}
func UnBase64(source string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(source)
}
