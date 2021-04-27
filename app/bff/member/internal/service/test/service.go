package test

import (
	"context"
	"go-kartos-study/app/bff/member/conf"
	"go-kartos-study/app/service/member/api/grpc"
	"go-kartos-study/pkg/net/netutil/breaker"
)

// Service .
type Service struct {
	c      *conf.Config
	breaker *breaker.Group
	memRPC grpc.MemberRPCClient
}

// New init service.
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c: c,
		breaker : breaker.NewGroup(nil),
	}
	var err error
	if s.memRPC, err = grpc.NewClient(c.MemberClient); err != nil {
		panic(err)
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
	//brk.MarkSuccess()
	// 正常情况下
	// doSomething
	//onBreaker(breaker breaker.Breaker, err *error) {
	//	if err != nil && *err != nil {
	//		breaker.MarkFailed()
	//	} else {
	//		breaker.MarkSuccess()
	//	}
	//}
	return
}
