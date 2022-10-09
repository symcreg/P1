package controller

import (
	"P1/config"
	"P1/utility"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func RegisterHandler(c *gin.Context) {
	var user config.User
	c.ShouldBindJSON(&user)                         //密码应为前端rsa公钥加密后的密文
	user.Password, _ = utility.UnRSA(user.Password) //使用私钥解密
	db, err := gorm.Open("sqlite3", "P1.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//密码应为6到16位，用户名应为4到16位
	if len(user.Password) < 6 || len(user.Password) > 16 || len(user.Username) < 4 || len(user.Username) > 16 {
		c.JSON(403, gin.H{
			"code": 403,
			"msg":  "wrong format",
		}) //判断是否符合规则
	} else {
		user.Password, _ = utility.PasswordHash(user.Password) //hash password
		db.Create(&user)
	}
}
func LoginHandler(c *gin.Context) {
	Id, _ := c.Get("Id")
	User, _ := c.Get("User")
	if Id != "" {
		c.JSON(200, gin.H{
			"code":     200,
			"msg":      "ReLogin",
			"uid":      Id,
			"username": User,
		})
		return
	} //有token，不必登录
	db, err := gorm.Open("sqlite3", "P1.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var user config.User
	var userFromDB config.User
	c.ShouldBindJSON(&user)                         //密码应为前端rsa公钥加密后的密文
	user.Password, _ = utility.UnRSA(user.Password) //使用私钥解密
	if err != nil {
		panic(err)
	}
	db.Where("username=?", user.Username).First(&userFromDB)
	if user.Username != userFromDB.Username {
		c.JSON(403, gin.H{
			"code": 403,
			"msg":  "wrong username or password",
		}) //用户名或密码错误
	} else {
		if utility.PasswordVerify(user.Password, userFromDB.Password) {
			c.JSON(200, gin.H{
				"code":     200,
				"msg":      "login success",
				"uid":      userFromDB.Id,
				"username": userFromDB.Username,
			}) //登录成功，前端应请求token
		} else {
			c.JSON(403, gin.H{
				"code": 403,
				"msg":  "wrong username or password",
			}) //用户名或密码错误
		}
	}
}
