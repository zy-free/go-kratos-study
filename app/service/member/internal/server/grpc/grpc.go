package grpc

import (
	"go-kartos-study/app/service/member/api/grpc"
	"go-kartos-study/app/service/member/internal/service"
	"go-kartos-study/pkg/net/rpc/warden"
)

type Server struct {
	svc *service.Service
}

// New
func New(c *warden.ServerConfig, svc *service.Service) *warden.Server {
	ws := warden.NewServer(c)

	grpc.RegisterMemberRPCServer(ws.Server(), &Server{svc: svc})
	ws, err := ws.Start()
	if err != nil {
		panic(err)
	}
	return ws
}
