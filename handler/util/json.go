package util

import (
	"crypto-market/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

//错误返回的json
func ResponseErrorJson(c *gin.Context, code int64) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  common.Errors[code],
	})
	c.Abort()
}

//
func ResponseErrorJsonWithMsg(c *gin.Context, code int64, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
	c.Abort()
}

//正确返回的json
func ResponseJson(c *gin.Context, msg, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  msg,
		"data": data,
	})
}

//
func ResponseErrorFormatJson(c *gin.Context, code int64, lang string, v ...interface{}) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": code,
		"msg":  common.Errors[code],
	})
}

//获取语言
// func GetLang(c *gin.Context) string {
// 	language := c.Request.Header.Get("Language")
// 	if len(language) > 0 {
// 		return language
// 	}
// 	return lang.ZH_CN
// }

//获取版本号
func GetVersion(c *gin.Context) string {
	version := c.Request.Header.Get("version")
	if len(version) > 0 {
		return version
	}
	return "1.1.0"
}

// 获取客户端
func GetClient(c *gin.Context) string {
	return c.Request.Header.Get("x-client")
}
