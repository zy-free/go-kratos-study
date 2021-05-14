package favorite

import (
	"context"
	"go-kartos-study/app/admin/member/conf"
	"go-kartos-study/app/admin/member/internal/dao/favorite"
	"go-kartos-study/app/admin/member/internal/model"
)

// Service .
type Service struct {
	c      *conf.Config
	favDao *favorite.Dao
}

// New init service.
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:      c,
		favDao: favorite.New(c),
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

func (s *Service) GetFavoriteByID(ctx context.Context, arg *model.GetFavoriteByIDReq) (favorite *model.Favorite, err error) {
	favorite, err = s.favDao.GetFavoriteByID(ctx, arg.Id)
	return
}
