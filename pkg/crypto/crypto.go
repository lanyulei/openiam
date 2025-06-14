package crypto

import (
	"encoding/base64"

	"github.com/lanyulei/toolkit/cipher"
	"github.com/lanyulei/toolkit/logger"
)

/*
  @Author : lanyulei
  @Desc :
*/

// AesEncryptCBC 加密
func AesEncryptCBC(key, value []byte) (result string, err error) {
	var (
		val []byte
	)

	key = key[3 : len(key)-3]
	val, err = cipher.AesEncryptCBC(key, value)
	if err != nil {
		logger.Errorf("AES encryption failed with error: %v", err)
		return
	}
	result = base64.StdEncoding.EncodeToString(val)
	return
}

// AesDecryptCBC 解密
func AesDecryptCBC(key []byte, value string) (result string, err error) {
	var (
		val    []byte
		aesVal []byte
	)
	key = key[3 : len(key)-3]
	val, err = base64.StdEncoding.DecodeString(value)
	if err != nil {
		logger.Errorf("base64 decoding failed with error: %v", err)
		return
	}
	aesVal, err = cipher.AesDecryptCBC(key, val)
	if err != nil {
		logger.Errorf("AES decryption failed with error: %v", err)
		return
	}

	result = string(aesVal)
	return
}
