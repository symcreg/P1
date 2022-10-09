package config

import (
	"time"
)

const TokenExpireDuration = time.Hour * 2 * 24 //过期时间2*24h

var Secret []byte = []byte("5211") //盐

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
