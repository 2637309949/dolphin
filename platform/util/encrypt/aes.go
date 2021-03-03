package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// AesEncrypt deifned
func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = pkcs5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key)
	encrypted := make([]byte, len(origData))
	blockMode.CryptBlocks(encrypted, origData)
	return []byte(base64.URLEncoding.EncodeToString(encrypted)), nil
}

// AesDecrypt defined
func AesDecrypt(crypted, key []byte) ([]byte, error) {
	decryptBytes, err := base64.URLEncoding.DecodeString(string(crypted))
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	decrypted := make([]byte, len(decryptBytes))
	blockMode.CryptBlocks(decrypted, crypted)
	decrypted = pkcs5UnPadding(decrypted)
	return decrypted, nil
}
