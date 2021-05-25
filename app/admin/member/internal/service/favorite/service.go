package favorite

import (
	"context"
	"go-kartos-study/app/admin/member/internal/dao/favorite"
	"go-kartos-study/app/admin/member/internal/model"
)

// Service .
type Service struct {
	favDao *favorite.Dao
}

// New init service.
func New(favDao *favorite.Dao) (s *Service) {
	s = &Service{
		favDao: favDao,
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
	favorite, err = s.favDao.GetFavoriteByID(ctx, arg.ID)
	return
}
