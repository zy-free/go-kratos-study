package http

import (
	"go-kartos-study/app/admin/member/conf"
	"go-kartos-study/app/admin/member/internal/service/favorite"
	"go-kartos-study/app/admin/member/internal/service/member"
	"go-kartos-study/pkg/log"
	bm "go-kartos-study/pkg/net/http/blademaster"
	nmd "go-kartos-study/pkg/net/metadata"
	"net/http"
)

var (
	favSvc  *favorite.Service
	memSvc  *member.Service
)

func initService(c *conf.Config) {
	favSvc = favorite.New(c)
	memSvc = member.New(c)
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
		c.Context = nmd.NewContext(c.Context, nmd.MD{nmd.Mid: "4", nmd.Color: "test"})
		c.Next()
	}
}

func route(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/x/admin", metadataMiddleware())
	{
		memGroup := g.Group("/members")
		{
			memGroup.GET("/info/getById", getMemberByID)
			//memGroup.GET("/info/getByPhone", getMemberByPhone)
			//memGroup.GET("/maxAge", getMemberMaxAge)
			//memGroup.GET("/queryByName", queryMemberByName)
			//memGroup.GET("/queryByIds", queryMemberByIDs)
			//memGroup.GET("/list", listMember)
			//memGroup.GET("/export", exportMember)
			memGroup.POST("", addMember)
			memGroup.POST("/batch", batchAddMember)
			memGroup.PUT("/update", updateMember)
			//memGroup.PUT("/set", setMember)
			//memGroup.PUT("/sort", sortMember)
			//memGroup.DELETE("/del", delMember)
		}

		favGroup := g.Group("/favorites")
		{
			favGroup.POST("/", addFavorite)
			favGroup.GET("/info/getById", getFavoriteByID)
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
