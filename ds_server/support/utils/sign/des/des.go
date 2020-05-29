package sign

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	consVal "ds_server/support/utils/constex"
)

func PKCS5Padding(srcData []byte, blockSize int) []byte {
	padding := blockSize - len(srcData)%8
	padtxt := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(srcData, padtxt...)
}

func PKCS5UnPadding(srcData []byte) []byte {
	length := len(srcData)
	unpadding := int(srcData[length-1])
	return srcData[:length-unpadding]
}

func DesEncrypt(srcData []byte, key []byte) []byte {
	block, _ := des.NewCipher(key)
	srcData = PKCS5Padding(srcData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(srcData))
	blockMode.CryptBlocks(crypted, srcData)
	return crypted
}

func DesDecrypt(srcData []byte, key []byte) []byte {
	block, _ := des.NewCipher(key)
	blockMode := cipher.NewCBCDecrypter(block, key)
	buf := make([]byte, len(srcData))
	blockMode.CryptBlocks(buf, srcData)
	buf = PKCS5UnPadding(buf)
	return buf
}

func EncryptPasswd(src string) string {
	ret := DesEncrypt([]byte(src), []byte(consVal.EncryptKey))//8
	return base64.StdEncoding.EncodeToString(ret)
}
func EncryptPasswdEx(src string,salt string) string {
	ret := DesEncrypt([]byte(src), []byte(salt))
	return base64.StdEncoding.EncodeToString(ret)
}
func DecryptPasswd(srcIn string) string {
	ret, _ := base64.StdEncoding.DecodeString(srcIn)
	src := DesDecrypt(ret, []byte(consVal.EncryptKey))//8
	return string(src)
}
//example:
//mysqlex := `root:dbjzcloudgame@(117.50.93.89:4040)/jz_cloud_game?charset=utf8&parseTime=true&loc=Asia%2FShanghai`
//ed := &EncryDecry{}
//src :=ed.EncryptPasswd(mysqlex)//加密后
//fmt.Println(src)
//ret := ed.DecryptPasswd(src)//解密后
//fmt.Println(ret)




