package encoder

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	b64 "encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func MD5(in string) string {
	alg := md5.New()
	alg.Write([]byte(in))
	return hex.EncodeToString(alg.Sum(nil))
}

func SHA256(in string) string {
	alg := sha256.New()
	alg.Write([]byte(in))
	return hex.EncodeToString(alg.Sum(nil))
}

func SHA512(in string) string {
	alg := sha512.New()
	alg.Write([]byte(in))
	return hex.EncodeToString(alg.Sum(nil))
}

func SHA1(in string) string {
	h := sha1.New()
	h.Write([]byte(in))
	return hex.EncodeToString(h.Sum(nil))
}

func Hmac(data, key string) string {
	hmac := hmac.New(md5.New, []byte(key))
	hmac.Write([]byte(data))
	return hex.EncodeToString(hmac.Sum([]byte("")))
}

func HmacToBase64(data, key string) string {
	hmac := hmac.New(sha1.New, []byte(key))
	hmac.Write([]byte(data))
	sEnc := b64.StdEncoding.EncodeToString([]byte(hmac.Sum(nil)))
	return sEnc
}

func Base64Encode(in []byte) string {
	return b64.StdEncoding.EncodeToString(in)
}

func Base64Decode(in string) []byte {
	out, err := b64.StdEncoding.DecodeString(in)
	if err != nil {
		return nil
	}
	return out
}

// instead of original unserialize function
func UnserializeFromJson(jsonstr string, st interface{}) error {
	d := json.NewDecoder(strings.NewReader(jsonstr))
	d.UseNumber()
	return d.Decode(st)
}

func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func RsaEncryptByPub(origData, publicKey []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecryptByPri(ciphertext, privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
