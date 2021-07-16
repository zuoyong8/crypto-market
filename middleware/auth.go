package middleware

//用户授权中间件
// func UserAuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		apiSecret := strings.Trim(c.Request.Header.Get(common.X_API_KEY), " ")
// 		apiUser := strings.Trim(c.Request.Header.Get(common.X_API_USER), " ")

// 		if apiSecret == "" || apiUser == "" || apiUser != key.Keys[apiSecret] {
// 			util.ResponseErrorJson(c, 1010012)
// 			return
// 		}
// 		c.Next()
// 	}
// }
