package test

import (
	"context"
	"fmt"
	"net/url"

	"go-kartos-study/app/service/member/api/grpc"
	"go-kartos-study/pkg/log"
	xhttp "go-kartos-study/pkg/net/http/blademaster"
	"go-kartos-study/pkg/net/metadata"
	"go-kartos-study/pkg/net/netutil/breaker"
)

// Service .
type Service struct {
	breaker    *breaker.Group
	memRPC     grpc.MemberRPCClient
	httpClient *xhttp.Client
}

// New init service.
func New(httpClient *xhttp.Client, memRPC grpc.MemberRPCClient) (s *Service) {
	s = &Service{
		breaker:    breaker.NewGroup(nil),
		httpClient: httpClient,
		memRPC:     memRPC,
	}
	return s
}

// Close service
func (s *Service) Close() {
}

// Ping service
func (s *Service) Ping(c context.Context) (err error) {
	return
}

func (s *Service) GrpcErrorTest(ctx context.Context) (err error) {
	_, err = s.memRPC.ErrorTest(ctx, &grpc.EmptyReq{})
	return
}

func (s *Service) MetadataErrorTest(ctx context.Context) (err error) {
	_, err = s.memRPC.MetadataTest(ctx, &grpc.EmptyReq{})
	return
}

func (s *Service) BreakerTest(ctx context.Context) (err error) {
	brk := s.breaker.Get("break_test")
	if err = brk.Allow(); err != nil {
		return
	}
	brk.MarkFailed()
	// brk.MarkSuccess()
	// 正常情况下
	// doSomething
	// onBreaker(breaker breaker.Breaker, err *error) {
	//	if err != nil && *err != nil {
	//		breaker.MarkFailed()
	//	} else {
	//		breaker.MarkSuccess()
	//	}
	// }
	return
}

func (s *Service) HTTPClientTest(ctx context.Context) (err error) {
	ip := metadata.String(ctx, metadata.RemoteIP)
	params := url.Values{}
	params.Set("id", "杭州")
	var res struct {
		Status     string
		City       string
		Citycode   string
		Weather36h []struct {
			Temp    string
			Time    string
			Weather string
		}
	}
	err = s.httpClient.Get(ctx, fmt.Sprintf("%s?%s", "http://api.help.bj.cn/apis/weather36h/", params.Encode()), ip, params, &res)
	if err != nil {
		return
	}
	log.Info("httpClient.Get res:%s", res)

	body := struct {
		Phone string `json:"phone"`
	}{Phone: "test1238139"}
	var res2 struct {
		Code    int
		Message string
	}
	err = s.httpClient.Post(ctx, "http://127.0.0.1:8000/x/bff/members", ip, body, &res2)
	if err != nil {
		return
	}
	log.Info("httpClient.Post res:%v", res2)

	return nil
}
