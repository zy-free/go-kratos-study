package service

import (
	"context"
	"go-kartos-study/app/service/member/conf"
	"go-kartos-study/app/service/member/internal/dao/favorite"
	"go-kartos-study/app/service/member/internal/dao/member"
)

// Service .
type Service struct {
	c      *conf.Config
	favDao *favorite.Dao
	memDao *member.Dao
}

// New init service.
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:      c,
		favDao: favorite.New(c),
		memDao: member.New(c),
	}
	return s
}

// Close service
func (s *Service) Close() {
	s.favDao.Close()
	s.memDao.Close()
}

// Ping service
func (s *Service) Ping(c context.Context) (err error) {
	s.favDao.Ping(c)
	s.memDao.Ping(c)
	return
}
