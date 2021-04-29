package http

import (
	"go-kartos-study/app/bff/member/conf"
	"go-kartos-study/app/bff/member/internal/service/favorite"
	"go-kartos-study/app/bff/member/internal/service/member"
	"go-kartos-study/app/bff/member/internal/service/test"
	"go-kartos-study/pkg/log"
	bm "go-kartos-study/pkg/net/http/blademaster"
	nmd "go-kartos-study/pkg/net/metadata"
	"net/http"
)

var (
	favSvc  *favorite.Service
	memSvc  *member.Service
	testSvc *test.Service
)

func initService(c *conf.Config) {
	favSvc = favorite.New(c)
	memSvc = member.New(c)
	testSvc = test.New(c)
}

//CloseService close all service
func CloseService() {
	favSvc.Close()
	memSvc.Close()
}

// Init http.
func Init(c *conf.Config) {
	initService(c)

	engine := bm.DefaultServer(c.HTTPServer)
	route(engine)
	if err := engine.Start(); err != nil {
		log.Error("engine.Start() error(%v)", err)
		panic(err)
	}
}

func metadataMiddleware() bm.HandlerFunc {
	return func(c *bm.Context) {
		// more metadata from gateway,this just test
		c.Context = nmd.NewContext(c.Context, nmd.MD{nmd.Mid: "4", nmd.Color: "green2"})
		c.Next()
	}
}

func route(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/x/bff", metadataMiddleware())
	{
		memGroup := g.Group("/members")
		{
			memGroup.GET("/info/getById", getMemberByID)
			memGroup.GET("/info/getByPhone", getMemberByPhone)
			memGroup.GET("/maxAge", getMemberMaxAge)
			memGroup.GET("/queryByName", queryMemberByName)
			memGroup.GET("/queryByIds", queryMemberByIDs)
			memGroup.GET("/list", listMember)
			memGroup.GET("/export", exportMember)
			memGroup.POST("", addMember)
			memGroup.POST("/init", initMember)
			memGroup.POST("/batch", batchAddMember)
			memGroup.PUT("/update", updateMember)
			memGroup.PUT("/set", setMember)
			memGroup.PUT("/sort", sortMember)
			memGroup.DELETE("/del", delMember)
		}

		favGroup := g.Group("/favorites")
		{
			favGroup.POST("/", addFavorite)
			favGroup.GET("/info/getById", getFavoriteByID)
		}

		testGroup := g.Group("/test")
		{
			testGroup.GET("/grpcError", grpcErrorTest)
			testGroup.GET("/maxConnTest", maxConnTest)
			testGroup.GET("/traceTest", traceTest)
			testGroup.GET("/metadataTest", metadataTest)
			testGroup.GET("/breakerTest", breakerTest)
			testGroup.GET("/errGroupTest", errGroupTest)
			testGroup.GET("/errGroupWithCancelTest", errGroupWithCancelTest)
			testGroup.GET("/goSafeTest", goSafeTest)
			testGroup.GET("/httpClientTest", httpClientTest)
			testGroup.GET("/cancelTest", cancelTest)
		}
	}
}

func ping(c *bm.Context) {
	if err := memSvc.Ping(c); err != nil {
		log.Error("member-service ping error(%v)", err)
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
	if err := favSvc.Ping(c); err != nil {
		log.Error("fav-service ping error(%v)", err)
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
	c.AbortWithStatus(http.StatusOK)
}
