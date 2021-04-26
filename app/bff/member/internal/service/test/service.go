package test

import (
	"context"
	"go-kartos-study/app/bff/member/conf"
	"go-kartos-study/app/service/member/api/grpc"
)

// Service .
type Service struct {
	c      *conf.Config
	memRPC grpc.MemberRPCClient
}

// New init service.
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c: c,
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