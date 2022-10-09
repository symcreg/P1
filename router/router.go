package router

import (
	"P1/utility/middleware"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func SetRouter() {
	Router.Use(middleware.Authorization) //token verify

}

/*
1xx：指示信息 - 表示请求已接收，继续处理
2xx：成功 - 表示请求已被成功接收、理解、接受
3xx：重定向 - 要完成请求必须进行更进一步的操作
4xx：客户端错误 - 请求有语法错误或请求无法实现
5xx：服务器端错误 - 服务器未能实现合法的请求
200： OK - 客户端请求成功
400： Bad Request - 客户端请求有语法错误，不能被服务器所理解
401： Unauthorized - 请求未经授权，这个状态代码必须和WWW-Authenticate报头域一起使用
403： Forbidden - 服务器收到请求，但是拒绝提供服务
404： Not Found - 请求资源不存在，eg：输入了错误的URL
500： Internal Server Error - 服务器发生不可预期的错误
503： Server Unavailable - 服务器当前不能处理客户端的请求，一段时间后,可能恢复正常
GET： 请求获取Request-URI所标识的资源
POST： 在Request-URI所标识的资源后增加新的数据
HEAD： 请求获取由Request-URI所标识的资源的响应消息报头
PUT： 请求服务器存储或修改一个资源，并用Request-URI作为其标识
DELETE： 请求服务器删除Request-URI所标识的资源
TRACE： 请求服务器回送收到的请求信息，主要用于测试或诊断
CONNECT： 保留将来使用
OPTIONS： 请求查询服务器的性能，或者查询与资源相关的选项和需求
*/
