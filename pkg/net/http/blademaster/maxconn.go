package blademaster

import (
	"go-kartos-study/pkg/log"
	"go-kartos-study/pkg/sync/limit"
	"net/http"
)

// MaxConns returns a middleware that limit the concurrent connections.
func MaxConn(n int) HandlerFunc {
	if n <= 0 {
		return func(c *Context) {
			c.Next()
		}
	}
	latch := limit.NewLimit(n)
	return func(c *Context) {
		if latch.TryBorrow() {
			defer func() {
				if err := latch.Return(); err != nil {
					log.Error("latch.Return error:%+v", err)
				}
			}()
			c.Next()
		} else {
			log.Error("concurrent connections over %d, rejected with code %d",n,http.StatusServiceUnavailable)
			c.AbortWithStatus(http.StatusServiceUnavailable)
			return
		}
	}

}
