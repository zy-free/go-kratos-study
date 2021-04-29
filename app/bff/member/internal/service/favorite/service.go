package favorite

import (
	"context"
	"go-kartos-study/app/bff/member/conf"
	"go-kartos-study/app/bff/member/internal/model"
	"go-kartos-study/app/service/member/api/grpc"
	"go-kartos-study/pkg/ecode"
	xtime "go-kartos-study/pkg/time"
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

func (s *Service) GetFavoriteByID(ctx context.Context, arg *model.GetFavoriteByIDReq) (result *model.GetFavoriteResp, err error) {
	resp, err := s.memRPC.GetFavoriteByID(ctx, &grpc.GetFavoriteByIDReq{
		Id: arg.Id,
	})
	if err != nil {
		return nil, ecode.ErrQuery
	}
	result = &model.GetFavoriteResp{Favorite: model.Favorite{
		Id:     resp.Id,
		Mid:    resp.Mid,
		Name:   resp.Name,
		HintAt: xtime.Time(resp.HintAt),
	}}
	return
}
