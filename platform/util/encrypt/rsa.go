package encrypt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"sort"
)

// SignWithPKCSV1SHA1 defined
func SignWithPKCSV1SHA1(playload map[string]string, privateKey string) (string, error) {
	pk, err := priKeyFromByte([]byte(privateKey))
	if err != nil {
		return "", err
	}
	mstr := map2SignUri(playload)
	hash := sha1.New()
	hash.Write([]byte(mstr))
	retByte, err := rsa.SignPKCS1v15(rand.Reader, pk, crypto.SHA1, hash.Sum(nil))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(retByte), err
}

// VerifyWithPKCSV1SHA1 defined
func VerifyWithPKCSV1SHA1(playload map[string]string, signed, publicKey string) error {
	pk, err := pubKeyFromByte([]byte(publicKey))
	if err != nil {
		return err
	}
	mstr := map2SignUri(playload)
	hash := sha1.New()
	hash.Write([]byte(mstr))
	descB64, err := base64.StdEncoding.DecodeString(signed)
	if err != nil {
		return err
	}
	return rsa.VerifyPKCS1v15(pk, crypto.SHA1, hash.Sum(nil), descB64)
}

func priKeyFromByte(privateKey []byte) (*rsa.PrivateKey, error) {
	pk := []byte(`
-----BEGIN RSA PRIVATE KEY-----
` + string(privateKey) + `
-----END RSA PRIVATE KEY-----
and some more`)

	block, _ := pem.Decode(pk)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		fmt.Println(block.Type)
		return nil, errors.New("failed to decode PEM block containing private key")
	}
	pub, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	return pub.(*rsa.PrivateKey), err
}

func pubKeyFromByte(publicKey []byte) (*rsa.PublicKey, error) {
	pk := []byte(`
-----BEGIN PUBLIC KEY-----
` + string(publicKey) + `
-----END PUBLIC KEY-----
and some more`)
	block, _ := pem.Decode(pk)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("failed to decode PEM block containing public key")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	return pub.(*rsa.PublicKey), err
}

func rsaDecrypt(key *rsa.PrivateKey, encrypted []byte) ([]byte, error) {
	return rsa.DecryptOAEP(sha256.New(), rand.Reader, key, encrypted, []byte{})
}

func rsaEncrypted(key *rsa.PublicKey, message []byte) ([]byte, error) {
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, key, message, []byte{})
}

func rsaSignPKCS1v15(key *rsa.PrivateKey, data []byte) (string, error) {
	hash := sha256.New()
	hash.Write(data)
	retByte, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, hash.Sum(nil))
	return base64.StdEncoding.EncodeToString(retByte), err
}

func rsaVerifyPKCS1v15(key *rsa.PublicKey, data []byte, base64Sign string) error {
	hash := sha256.New()
	hash.Write(data)
	descB64, err := base64.StdEncoding.DecodeString(base64Sign)
	err = rsa.VerifyPKCS1v15(key, crypto.SHA256, hash.Sum(nil), descB64)
	return err
}

func map2SignUri(puData interface{}) string {
	var puJSON map[string]string
	var puKeys = make([]string, 0, len(puJSON))
	puByte, _ := json.Marshal(puData)
	json.Unmarshal(puByte, &puJSON)
	for k := range puJSON {
		if k != "sign" {
			puKeys = append(puKeys, k)
		}
	}
	sort.Strings(puKeys)
	var signString = ""
	for _, k := range puKeys {
		if signString != "" {
			signString = signString + "&" + k + "=" + puJSON[k]
		} else {
			signString = signString + k + "=" + puJSON[k]
		}
	}
	return signString
}
