package utility

import (
	"P1/config"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Jwt struct {
	Id                 int
	User               string //用户名称
	jwt.StandardClaims        //jwt标准claims
}

func GenToken(claim Jwt) (string, error) {
	claim.ExpiresAt = time.Now().Add(config.TokenExpireDuration).Unix() //过期时间
	claim.Issuer = "SYMC"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	SignedToken, err := token.SignedString(config.Secret)
	if err != nil {
		return "", err
	}
	return SignedToken, err
}
func ParseToken(tokenStr string) (*Jwt, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Jwt{}, func(token *jwt.Token) (interface{}, error) {
		return config.Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claim, ok := token.Claims.(*Jwt); ok && token.Valid {
		return claim, nil
	}
	return nil, errors.New("invalid token")
}
func GetTokenHandler(c *gin.Context) {
	var user config.User
	var claim Jwt
	c.ShouldBindJSON(&user)
	claim.Id = user.Id
	claim.User = user.Username
	token, err := GenToken(claim)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "error",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"msg":   "success",
		"token": token,
	})
	return
}
