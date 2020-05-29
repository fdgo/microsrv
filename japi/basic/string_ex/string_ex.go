package string_ex

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
	"unicode/utf8"
	"unsafe"
)
//使用 utf8.RuneCountInString()统计字符串长度
func Length(str string) int {
	return utf8.RuneCountInString(str)
}
func SubString(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}
//获取一个Guid
func GetGuid() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return Md5(Base64(b))
}
func Md5(s string) string {
	h := md5.New()
	h.Write(StringToSliceByte(s))
	return hex.EncodeToString(h.Sum(nil))
}
// string()
func SliceByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// []byte()
func StringToSliceByte(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
func Base64(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

