package aes

import (
	"bytes"
	"crypto/aes"
)

// Aes Tool
type Aes struct {
	Key       []byte
	BlockSize int
}

// ECBEncrypt 加密码
func (a *Aes) ECBEncrypt(src []byte) ([]byte, error) {
	block, err := aes.NewCipher(a.Key)
	if err != nil {
		return nil, err
	}

	src = a.padding(src)
	encryptData := make([]byte, len(src))
	tmpData := make([]byte, a.BlockSize)

	for index := 0; index < len(src); index += a.BlockSize {
		block.Encrypt(tmpData, src[index:index+a.BlockSize])
		copy(encryptData, tmpData)
	}
	return encryptData, nil
}

// ECBDecrypt 解密
func (a *Aes) ECBDecrypt(src []byte) ([]byte, error) {
	block, err := aes.NewCipher(a.Key)
	if err != nil {
		return nil, err
	}
	decryptData := make([]byte, len(src))
	tmpData := make([]byte, a.BlockSize)
	for index := 0; index < len(src); index += a.BlockSize {
		block.Decrypt(tmpData, src[index:index+a.BlockSize])
		copy(decryptData, tmpData)
	}

	dst := a.unPandding(decryptData)
	return dst, nil
}

func (a *Aes) padding(src []byte) []byte {
	paddingCount := aes.BlockSize - len(src)%aes.BlockSize
	if paddingCount == 0 {
		return src
	} else {
		return append(src, bytes.Repeat([]byte{byte(0)}, paddingCount)...)
	}
}

func (a *Aes) unPandding(dst []byte) []byte {
	for i := len(dst) - 1; ; i-- {
		if dst[i] != 0 {
			return dst[:i+1]
		}
	}
}
