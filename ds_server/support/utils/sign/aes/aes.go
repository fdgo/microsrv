package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

//DES  密钥必须8位
//3DES 密钥必须24位
//AES  明文分组长度为128位（16字节），密钥长度可以为128位（16字节），192位（24字节），256位（32字节）

//DES AES 对称加密
//RSA  非对称加密

//在加密和解密之前，首先需要补码和去码操作

//补码   PKCS7Padding分组是以1-255为单位
func PKCS7Padding(srcData []byte, blockSize int) []byte {
	padding := blockSize - len(srcData)%blockSize
	padtxt := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(srcData, padtxt...)
}

//去码
func PKCS7UnPadding(srcData []byte) []byte {
	length := len(srcData)
	unpadding := int(srcData[length-1])
	return srcData[:length-unpadding]
}

//AES加密,加密会用到补码
func AesEncrypt(srcData []byte, key []byte) []byte {
	//首先检验密钥是否合法
	block, _ := aes.NewCipher(key)
	//补码
	srcData = PKCS7Padding(srcData, block.BlockSize())
	//设置加密方式
	blockMode := cipher.NewCBCEncrypter(block, key)
	//加密处理
	crypted := make([]byte, len(srcData))

	blockMode.CryptBlocks(crypted, srcData)

	return crypted

}

//AES解密,解密要用到去码
func AesDecrypt(srcData []byte, key []byte) []byte {
	//首先检验密钥是否合法
	block, _ := aes.NewCipher(key)
	//设置解码方式
	blockMode := cipher.NewCBCDecrypter(block, key)
	//创建缓冲区
	buf := make([]byte, len(srcData))
	//开始解密
	blockMode.CryptBlocks(buf, srcData)
	//去掉编码
	buf = PKCS7UnPadding(buf)

	return buf
}

////加密
//src := AesEncrypt([]byte("kongyixueyuan"),[]byte("1234567890123456"))//必须16位
////通过base64处理密文
//fmt.Println(base64.StdEncoding.EncodeToString(src))
//
////解密
//src2 :=AesDecrypt(src,[]byte("1234567890123456"))
//fmt.Println(string(src2))
