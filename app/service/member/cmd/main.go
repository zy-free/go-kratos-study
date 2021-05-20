package main

import (
	"flag"
	"github.com/davecgh/go-spew/spew"
	"go-kartos-study/app/service/member/appid"
	"go-kartos-study/app/service/member/conf"
	favoriteDao "go-kartos-study/app/service/member/internal/dao/favorite"
	memberDao "go-kartos-study/app/service/member/internal/dao/member"
	"go-kartos-study/app/service/member/internal/server/grpc"
	"go-kartos-study/app/service/member/internal/service"
	"go-kartos-study/pkg/cache/redis"
	"go-kartos-study/pkg/conf/env"
	"go-kartos-study/pkg/database/sql"
	"go-kartos-study/pkg/naming/etcd"
	"go-kartos-study/pkg/net/trace"
	"go-kartos-study/pkg/net/trace/jaeger"
	"go-kartos-study/pkg/queue/kafka"
	"go-kartos-study/pkg/sync/pipeline"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"go-kartos-study/pkg/log"
)


func initApp(c *conf.Config) (closeFunc func()){
	db := sql.NewMySQL(c.Mysql)
	redis := redis.NewPool(c.Redis)
	publisher := kafka.NewPublisher(c.KafkaPublish)
	merge := pipeline.NewPipeline(c.Merge)

	favDao := favoriteDao.New(db)
	memDao := memberDao.New(db, redis, publisher)
	svc := service.New(favDao, memDao, merge)

	grpc.New(conf.Conf.GRPCServer, svc)
	return svc.Close
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
	//zipkin.Init(&zipkin.Config{
	//	Endpoint: "http://8.131.78.197:9433/api/v2/spans",
	//})
	jaeger.Init()
	defer trace.Close()

	closeFunc := initApp(conf.Conf)

	cancel := etcd.ETCDRegist(conf.Conf.ETCDConfig, appid.AppID, strings.Split(conf.Conf.GRPCServer.Addr, ":")[1], env.Color)
	defer cancel()
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
