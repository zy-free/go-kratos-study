package grpc

import (
	"context"
	"go-kartos-study/app/service/member/api/grpc"
)

func (s *Server) AddFavorite(c context.Context, req *grpc.AddFavoriteReq) (resp *grpc.IDResp, err error) {
	//s.dao.GetMemberByID()
	return
}

func (s *Server) GetFavoriteByID(ctx context.Context, req *grpc.GetFavoriteByIDReq) (resp *grpc.FavoriteResp, err error) {
	fav, err := s.svc.GetFavoriteByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	resp = &grpc.FavoriteResp{

		Id:     fav.Id,
		Mid:    fav.Mid,
		Name:   fav.Name,
		HintAt: fav.HintAt.Time().Unix(),
	}

	return
}
