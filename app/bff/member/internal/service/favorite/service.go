package favorite

import (
	"context"
	"go-kartos-study/app/bff/member/conf"
)

// Service .
type Service struct {
	c           *conf.Config
	//dao         *bnj.Dao
}

// New init service.
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:        c,
		//dao:      bnj.New(c),
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
