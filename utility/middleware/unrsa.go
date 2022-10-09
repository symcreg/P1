package middleware

import (
	"P1/utility"
	"crypto"
	"crypto/rsa"
	"github.com/gin-gonic/gin"
)

func UnRSA(c *gin.Context) { //根据私钥解密
	PrivateKey := utility.PrivateKey
	EncryptedData, _ := c.Get("EncryptedData")
	if EncryptedData == "" {
		c.Abort()
		return
	}
	DecryptedData, err := PrivateKey.Decrypt(nil, EncryptedData.([]byte), &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		panic("decrypt data error")
		c.Abort()
		return
	}
	c.Set("DecryptedData", DecryptedData)
	c.Next()
}
