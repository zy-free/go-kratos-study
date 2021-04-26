package service

import (
	"context"
	"go-kartos-study/app/service/member/internal/model"
)

func (s *Service) GetMemberByID(ctx context.Context,id int64) (member *model.Member, err error) {
	return s.memDao.GetMemberByID(ctx,id)
}

func (s *Service) GetMemberByPhone(ctx context.Context,phone string) (member *model.Member, err error) {
	return s.memDao.GetMemberByPhone(ctx,phone)
}

func (s *Service) GetMemberMaxAge(ctx context.Context) (age int64, err error) {
	return s.memDao.GetMemberMaxAge(ctx)
}

func (s *Service) CountMember(ctx context.Context) (count int64, err error) {
	return s.memDao.CountMember(ctx)
}

func (s *Service) ListMember(ctx context.Context) (members []*model.Member, err error) {
	return s.memDao.ListMember(ctx)
}

func (s *Service) QueryMemberByName(ctx context.Context,name string) (members []*model.Member, err error) {
	return s.memDao.QueryMemberByName(ctx,name)
}

func (s *Service) QueryMemberByIDs(ctx context.Context,ids []int64) (res map[int64]*model.Member, err error) {
	return s.memDao.QueryMemberByIDs(ctx,ids)
}

func (s *Service) AddMember(ctx context.Context,member *model.Member) (id int64, err error) {
	return s.memDao.AddMember(ctx,member)
}

func (s *Service) BatchAddMember(ctx context.Context,members []*model.Member) (affectRow int64, err error) {
	return s.memDao.BatchAddMember(ctx,members)
}

func (s *Service) InitMember(ctx context.Context,member *model.Member) (err error) {
	return s.memDao.InitMember(ctx,member)
}

func (s *Service) UpdateMember(ctx context.Context,member *model.Member) (err error) {
	return s.memDao.UpdateMember(ctx,member)
}

func (s *Service) SetMember(ctx context.Context,member *model.Member) (err error) {
	return s.memDao.SetMember(ctx,member)
}

func (s *Service) SortMember(ctx context.Context,args model.ArgMemberSort) (err error) {
	return s.memDao.SortMember(ctx,args)
}

func (s *Service) DelMember(ctx context.Context,id int64) (err error) {
	return s.memDao.DeleteMember(ctx,id)
}
