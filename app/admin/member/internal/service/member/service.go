package member

import (
	"context"
	"go-kartos-study/app/admin/member/internal/dao/member"
	"go-kartos-study/app/admin/member/internal/model"
)

// Service .
type Service struct {
	memDao *member.Dao
}

// New init service.
func New(memDao *member.Dao) (s *Service) {
	s = &Service{
		memDao: memDao,
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

func (s *Service) GetMemberByID(ctx context.Context, id int64) (member *model.Member, err error) {
	return s.memDao.GetMemberByID(ctx, id)
}

func (s *Service) GetMemberByPhone(ctx context.Context, phone string) (member *model.Member, err error) {
	return
	//return s.memDao.GetMemberByPhone(ctx, phone)
}

func (s *Service) GetMemberMaxAge(ctx context.Context) (age int64, err error) {
	return
	//return s.memDao.GetMemberMaxAge(ctx)
}

func (s *Service) CountMember(ctx context.Context) (count int64, err error) {
	return
	//return s.memDao.CountMember(ctx)
}

func (s *Service) ListMember(ctx context.Context) (members []*model.Member, err error) {
	return

	//return s.memDao.ListMember(ctx)
}

func (s *Service) QueryMemberByName(ctx context.Context, name string) (members []*model.Member, err error) {
	return
	//return s.memDao.QueryMemberByName(ctx, name)
}

func (s *Service) QueryMemberByIDs(ctx context.Context, ids []int64) (res map[int64]*model.Member, err error) {
	return
	//return s.memDao.QueryMemberByIDs(ctx, ids)
}

func (s *Service) AddMember(ctx context.Context, arg *model.AddMemberReq) (id int64, err error) {
	return s.memDao.AddMember(ctx,arg)
}

func (s *Service) BatchAddMember(ctx context.Context, args []*model.AddMemberReq) (affectRow int64, err error) {
	return s.memDao.BatchAddMember(ctx, args)
}

func (s *Service) UpdateMember(ctx context.Context, arg *model.UpdateMemberReq) (err error) {
	return s.memDao.UpdateMember(ctx, arg)
}

// attr demo
func (s *Service) UpdateMemberLock(ctx context.Context, id int64, lock int32) (err error) {
	return
	//return s.memDao.UpdateMemberAttr(ctx, id, mem.Attr)
}

func (s *Service) SetMember(ctx context.Context, member *model.Member) (err error) {
	return
	//return s.memDao.SetMember(ctx, member)
}

func (s *Service) SortMember(ctx context.Context, args model.ArgMemberSort) (err error) {
	return
	//return s.memDao.SortMember(ctx, args)
}

func (s *Service) DelMember(ctx context.Context, id int64) (err error) {
	return
	//return s.memDao.DeleteMember(ctx, id)
}
