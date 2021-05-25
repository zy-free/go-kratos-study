package service

import (
	"context"
	"go-kartos-study/app/service/member/internal/dao/favorite"
	"go-kartos-study/app/service/member/internal/dao/member"
	"go-kartos-study/pkg/sync/pipeline"
)

// Service .
type Service struct {
	favDao *favorite.Dao
	memDao *member.Dao
	merge  *pipeline.Pipeline
}

// New init service.
func New(favDao *favorite.Dao, memDao *member.Dao, merge *pipeline.Pipeline) (s *Service) {
	s = &Service{
		favDao: favDao,
		memDao: memDao,
		merge:  merge,
	}
	s.initMerge()
	return s
}

// Close service
func (s *Service) Close() {
	s.favDao.Close()
	s.memDao.Close()
}

// Ping service
func (s *Service) Ping(c context.Context) (err error) {
	_ = s.favDao.Ping(c)
	_ = s.memDao.Ping(c)
	return
}
