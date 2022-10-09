package middleware

import (
	"P1/utility"
	"github.com/gin-gonic/gin"
)

func Authorization(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(401, gin.H{
			"code": 401,
			"msg":  "none token",
		})
		c.Abort()
		return
	}
	claim, err := utility.ParseToken(token)
	if err != nil {
		c.JSON(401, gin.H{
			"code": 401,
			"msg":  "invalid token",
		})
		c.Abort()
		return
	}
	c.Set("Uid", claim.Uid)
	c.Set("User", claim.User)
	c.Next()
}
