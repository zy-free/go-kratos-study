package service

import (
	"context"
	"go-kartos-study/app/service/member/internal/model"
	"go-kartos-study/pkg/log"
)

func (s *Service) GetFavoriteByID(ctx context.Context,id int64) (favorite *model.Favorite, err error) {
	favorite,err =  s.favDao.GetFavoriteByID(ctx,id)
	log.Info("fav:(%v)",favorite)
	return
}
