package main

import (
	"P1/db"
	"P1/router"
)

func main() {
	db.InitDB()                //初始化数据库
	router.SetRouter()         //设置路由
	router.Router.Run(":8080") //启动监听
}
