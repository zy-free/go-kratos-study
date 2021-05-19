package main

import (
	"flag"
	"github.com/davecgh/go-spew/spew"
	"go-kartos-study/app/bff/member/internal/server/http"
	"go-kartos-study/app/bff/member/internal/service/favorite"
	"go-kartos-study/app/bff/member/internal/service/member"
	"go-kartos-study/app/bff/member/internal/service/test"
	"go-kartos-study/app/service/member/api/grpc"
	xhttp "go-kartos-study/pkg/net/http/blademaster"
	"go-kartos-study/pkg/net/trace/zipkin"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-kartos-study/app/bff/member/conf"
	"go-kartos-study/pkg/log"
)

func initApp(c *conf.Config) (closeFunc func()) {
	httpClient := xhttp.NewClient(c.HTTPClient)
	memRPC, err := grpc.NewClient(c.MemberClient)
	if err != nil {
		panic(err)
	}

	testService := test.New(httpClient, memRPC)
	memService := member.New(memRPC)
	favService := favorite.New(memRPC)

	err = http.Init(c.HTTPServer, favService, memService, testService)
	if err != nil{
		panic(err)
	}
	return  http.CloseService
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
	//trace.Init(conf.Conf.Tracer)
	//jaeger.Init()
	zipkin.Init(&zipkin.Config{
		Endpoint: "http://8.131.78.197:9411/api/v2/spans",
	})
	//defer trace.Close()

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


