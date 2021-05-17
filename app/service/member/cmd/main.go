package main

import (
	"flag"
	"github.com/davecgh/go-spew/spew"
	"go-kartos-study/app/service/member/appid"
	"go-kartos-study/app/service/member/conf"
	"go-kartos-study/pkg/conf/env"
	"go-kartos-study/pkg/naming/etcd"
	"go-kartos-study/pkg/net/trace"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"go-kartos-study/pkg/log"
)

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
