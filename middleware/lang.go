package middleware

import (
	"github.com/gin-gonic/gin"
)

//语言
func Language() gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.Request.Header.Get("Language")
		if lang != "" {
			c.Writer.Header().Set("Language", lang)
		}
		c.Next()
	}
}
