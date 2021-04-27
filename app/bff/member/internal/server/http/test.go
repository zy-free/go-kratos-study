package http

import (
	bm "go-kartos-study/pkg/net/http/blademaster"
	"go-kartos-study/pkg/net/trace"
	"time"
)

// 验证grpc的error透传,直接将grpc的server端错误截获
func grpcErrorTest(c *bm.Context) {
	c.JSON(nil, testSvc.GrpcErrorTest(c))
}

// 验证maxConn中间件，当前配置为1，压测10qps，则应为90%报错，生产建议10000 Conn
func maxConnTest(c *bm.Context) {
	time.Sleep(time.Second * 3)
	c.JSON("maxConn test", nil)
}

// 验证从header头中传入，如上层网关已经产生了trace
func traceTest(c *bm.Context) {
	t, _ := trace.FromContext(c.Context)
	c.JSON(t.TraceID(), nil)
}

func metadataTest(c *bm.Context) {
	c.JSON(nil, testSvc.MetadataErrorTest(c))
}

func breakerTest(c *bm.Context) {
	c.JSON(nil, testSvc.BreakerTest(c))
}
