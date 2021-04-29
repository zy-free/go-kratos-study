package blademaster

import (
	"go-kartos-study/pkg/ecode"
	"go-kartos-study/pkg/log"
	"go-kartos-study/pkg/net/netutil/breaker"
	"net/http"
)

func Breaker() HandlerFunc {

	group := breaker.NewGroup(nil)

	return func(c *Context) {
		brk := group.Get(c.method + c.RoutePath)
		if err := brk.Allow(); err != nil {
			log.Error("blademaster beaker ejected with code %d", http.StatusServiceUnavailable)
			c.AbortWithStatus(http.StatusServiceUnavailable)
			return
		}
		defer func() {
			err := c.Error
			if ecode.Cause(err) != ecode.OK {
				brk.MarkFailed()
			} else {
				brk.MarkSuccess()
			}
		}()
		c.Next()
	}

}
