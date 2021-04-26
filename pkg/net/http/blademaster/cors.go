package blademaster

import (
	"net/http"
)

// CORS returns the location middleware with default configuration.
func CORS() HandlerFunc {
	return func(c *Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		hander := c.Writer.Header()
		hander.Set("Access-Control-Allow-Origin", origin)
		hander.Set("Access-Control-Allow-Headers", "Content-Type,X-Token")
		hander.Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE,PUT")
		hander.Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		hander.Set("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
