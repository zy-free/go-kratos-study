package blademaster

import (
	"go-kartos-study/pkg/log"
	"net/http"
)

// MaxBytesHandler returns a middleware that limit reading of http request body.
func MaxByte(n int64) HandlerFunc {
	if n <= 0 {
		return func(c *Context)  {
			c.Next()
		}
	}

	return func(c *Context)  {
		if c.Request.ContentLength > n {
			log.Error("request entity too large, limit is %d, but got %d, rejected with code %d",n,c.Request.ContentLength,http.StatusRequestEntityTooLarge)
			c.AbortWithStatus(http.StatusRequestEntityTooLarge)
			return
		}else{
			c.Next()
		}
	}

}
