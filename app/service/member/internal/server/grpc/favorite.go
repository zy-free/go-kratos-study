package grpc

import (
	"context"
	"go-kartos-study/app/service/member/api/grpc"
)

func (s *Server) AddFavorite (c context.Context, req *grpc.AddFavoriteReq)  (resp *grpc.IDResp, err error) {
	//s.dao.GetMemberByID()
	return
}
