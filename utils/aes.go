package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// AES/CBC解密数据--PCKS#7填充,数据加密
func AesCBCEncrypt(source string, key string, iv string) (string, error) {
	// 生成16进制加密key
	newKey := []byte(key)
	block, err := aes.NewCipher(newKey)
	if err != nil {
		return "", err
	}
	// 数据处理
	dataLen := len([]byte(source))
	m := dataLen % 16
	if m != 0 {
		for i := 0; i < 16-m; i++ {
			source = source + " "
		}
	}
	newByte := []byte(source)
	// 初始向量IV必须是唯一
	newIV := []byte(iv)
	// block大小和初始向量大小一定要一致
	mode := cipher.NewCBCEncrypter(block, newIV)
	encryptData := make([]byte, len(newByte))
	// 加PKCS7填充
	content := PKCS7Padding(newByte, block.BlockSize())
	mode.CryptBlocks(encryptData, content)
	return base64.StdEncoding.EncodeToString(encryptData), nil
}

// AES/CBC解密数据--填充,数据解密 base64 string
func AesCBCDecrypt(source string, key string, iv string) (string, error) {
	// 生成16进制加密key
	newKey := []byte(key)
	block, err := aes.NewCipher(newKey)
	if err != nil {
		return "", err
	}
	// 16进制转换
	decodeBytes, err := base64.StdEncoding.DecodeString(source)
	if err != nil {
		return "", err
	}
	newIV := []byte(iv)
	mode := cipher.NewCBCDecrypter(block, newIV)
	retData := make([]byte, len(decodeBytes))
	mode.CryptBlocks(retData, decodeBytes)
	// 解PKCS7填充
	retData = PKCS7UnPadding(retData)
	return string(retData), nil
}

// PKCS7加填充/和PKCS5填充一样,只是填充字段多少的区别
func PKCS7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

// PKCS7解填充/和PKCS5填充一样,只是填充字段多少的区别
func PKCS7UnPadding(encrypt []byte) []byte {
	length := len(encrypt)
	unPadding := int(encrypt[length-1])
	return encrypt[:(length - unPadding)]
}
