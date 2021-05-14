package main

import (
	"flag"
	"github.com/davecgh/go-spew/spew"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-kartos-study/app/admin/member/conf"
	"go-kartos-study/app/admin/member/internal/server/http"
	"go-kartos-study/pkg/log"
	"go-kartos-study/pkg/net/trace"
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

	http.Init(conf.Conf)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("member-service get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			http.CloseService()
			log.Info("member-service exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}


