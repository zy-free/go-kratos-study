package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/davecgh/go-spew/spew"

	"go-kartos-study/app/admin/member/conf"
	favDao "go-kartos-study/app/admin/member/internal/dao/favorite"
	memDao "go-kartos-study/app/admin/member/internal/dao/member"
	"go-kartos-study/app/admin/member/internal/server/http"
	"go-kartos-study/app/admin/member/internal/service/favorite"
	"go-kartos-study/app/admin/member/internal/service/member"
	"go-kartos-study/pkg/database/orm"
	"go-kartos-study/pkg/log"
	"go-kartos-study/pkg/net/trace"
)

func initApp(c *conf.Config) (closeFunc func()) {
	db := orm.NewMySQL(c.ORM)
	db.LogMode(true)

	favDao := favDao.New(db)
	memDao := memDao.New(db)

	memService := member.New(memDao)
	favService := favorite.New(favDao)
	err := http.Init(c.HTTPServer, favService, memService)
	if err != nil {
		panic(err)
	}
	return http.CloseService
}

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Error("conf.Init() error(%v)", err)
		panic(err)
	}
	spew.Dump(conf.Conf)

	log.Init(conf.Conf.Log)
	defer log.Close()
	trace.Init(conf.Conf.Tracer)
	defer trace.Close()

	closeFunc := initApp(conf.Conf)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("member-service get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeFunc()
			log.Info("member-service exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
