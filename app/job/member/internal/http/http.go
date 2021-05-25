package http

import (
	"go-kartos-study/pkg/log"
	bm "go-kartos-study/pkg/net/http/blademaster"
	"net/http"
)

// Init http.
func Init(c *bm.ServerConfig) (err error) {
	engine := bm.DefaultServer(c)
	route(engine)
	if err := engine.Start(); err != nil {
		log.Error("engine.Start() error(%v)", err)
		return err
	}
	return nil
}

func route(e *bm.Engine) {
	e.Ping(ping)
}

func ping(c *bm.Context) {
	c.AbortWithStatus(http.StatusOK)
}
