package main

import (
	"go-kartos-study/app/bff/member/conf"
	"go-kartos-study/app/bff/member/internal/server/http"
	"go-kartos-study/app/bff/member/internal/service/favorite"
	"go-kartos-study/app/bff/member/internal/service/member"
	"go-kartos-study/app/bff/member/internal/service/test"
	"go-kartos-study/app/service/member/api/grpc"
	xhttp "go-kartos-study/pkg/net/http/blademaster"
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
