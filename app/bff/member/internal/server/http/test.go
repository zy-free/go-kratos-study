package http

import (
	"context"
	"fmt"
	"time"

	"go-kartos-study/pkg/ecode"
	bm "go-kartos-study/pkg/net/http/blademaster"
	"go-kartos-study/pkg/net/trace"
	"go-kartos-study/pkg/sync/errgroup"
	"go-kartos-study/pkg/sync/goroutine"
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

func errGroupTest(c *bm.Context) {
	g := errgroup.WithContext(c)

	// 总时长应该为2s
	g.Go(func(ctx context.Context) error {
		time.Sleep(time.Second)
		fmt.Println("errgroup go time sleep 1")
		return nil
	})

	g.Go(func(ctx context.Context) error {
		time.Sleep(time.Second * 2)
		fmt.Println("errgroup go time sleep 2")
		return nil
	})

	_ = g.Wait()
	c.JSON("ok", nil)
}

// 使用 WithCancel 时如果有一个人任务失败会导致所有*未进行或进行中*的任务被 cancel
func errGroupWithCancelTest(c *bm.Context) {
	var err error
	g := errgroup.WithCancel(c)

	g.Go(func(ctx context.Context) error {
		return ecode.ErrQuery
	})

	g.Go(func(ctx context.Context) error {
		time.Sleep(time.Second)
		fmt.Println("errgroup go time sleep 1")
		return nil
	})

	g.Go(func(ctx context.Context) error {
		<-ctx.Done()
		err = ctx.Err()
		return err
	})

	err = g.Wait()
	c.JSON("ok", err)
}

func goSafeTest(c *bm.Context) {
	goroutine.GoSafe(func() {
		panic("test panic")
	})
	c.JSON("ok", nil)
}

func httpClientTest(c *bm.Context) {
	c.JSON(nil, testSvc.HTTPClientTest(c))
}

func cancelTest(c *bm.Context) {
	ctx, cancel := context.WithCancel(c)
	cancel()
	c.JSON(nil, testSvc.GrpcErrorTest(ctx))
}
