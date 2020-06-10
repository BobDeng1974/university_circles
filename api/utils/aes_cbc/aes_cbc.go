package aes_cbc

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

/*CBC加密 按照golang标准库的例子代码
不过里面没有填充的部分,所以补上
*/

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func padding(src []byte) []byte {
	//填充个数
	paddingCount := aes.BlockSize - len(src)%aes.BlockSize
	if paddingCount == 0 {
		return src
	} else {
		//填充数据
		return append(src, bytes.Repeat([]byte{byte(0)}, paddingCount)...)
	}
}

//aes加密，填充秘钥key的16位，24,32分别对应AES-128, AES-192, or AES-256.
func AesCBCEncrypt(rawData, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	//填充原文
	blockSize := block.BlockSize()
	rawData = padding(rawData)
	//初始向量IV必须是唯一，但不需要保密
	cipherText := make([]byte, blockSize+len(rawData))
	//block大小 16
	// iv := cipherText[:blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	//block大小和初始向量大小一定要一致
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[blockSize:], rawData)

	return cipherText, nil
}

func AesCBCDncrypt(decodeStr string, key string, iv string) ([]byte, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(decodeStr)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	mode := cipher.NewCBCDecrypter(block, []byte(iv))

	resultData := make([]byte, len(decodeBytes))

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(resultData, decodeBytes)
	return resultData, nil
}

// func Dncrypt(rawData string, key []byte, iv []byte) (string, error) {
// 	data, err := base64.StdEncoding.DecodeString(rawData)
// 	if err != nil {
// 		return "", err
// 	}
// 	dnData, err := AesCBCDncryptt(data, key, iv)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(dnData), nil
// }

// func AesCBCDncryptt(encryptData, key []byte, iv []byte) ([]byte, error) {
// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		panic(err)
// 	}

// 	blockSize := block.BlockSize()

// 	if len(encryptData) < blockSize {
// 		panic("ciphertext too short")
// 	}
// 	// iv := encryptData[:blockSize]
// 	encryptData = encryptData[blockSize:]

// 	// CBC mode always works in whole blocks.
// 	if len(encryptData)%blockSize != 0 {
// 		panic("ciphertext is not a multiple of the block size")
// 	}

// 	mode := cipher.NewCBCDecrypter(block, iv)

// 	// CryptBlocks can work in-place if the two arguments are the same.
// 	mode.CryptBlocks(encryptData, encryptData)
// 	//解填充
// 	encryptData = PKCS7UnPadding(encryptData)
// 	return encryptData, nil
// }
