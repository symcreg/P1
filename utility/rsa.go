package utility

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

func genRSA() (*rsa.PrivateKey, rsa.PublicKey, error) {
	PrivateKey, err := rsa.GenerateKey(rand.Reader, 2048) //私钥，用此私钥解密
	if err != nil {
		return nil, rsa.PublicKey{}, err
	}
	PublicKey := PrivateKey.PublicKey //公钥，发送给前端以加密
	return PrivateKey, PublicKey, nil
}
func enRSA(data string, PublicKey rsa.PublicKey) (string, error) { //根据公钥加密
	EncryptedData, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&PublicKey,
		[]byte(data), //需加密的字符串
		nil)
	if err != nil {
		return "", err
	}
	return string(EncryptedData), nil
}
func unRSA(EncryptedData string, PrivateKey rsa.PrivateKey) (string, error) { //根据私钥解密
	DecryptedData, err := PrivateKey.Decrypt(nil, []byte(EncryptedData), &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		return "", err
	}
	return string(DecryptedData), nil
}
