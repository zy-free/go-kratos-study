package service

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"go-kartos-study/app/service/member/internal/model"
	"go-kartos-study/pkg/sync/pipeline"
	"strconv"
)

func (s *Service) GetMemberByID(ctx context.Context, id int64) (member *model.Member, err error) {
	// test
	s.addMerge(ctx, 1, 1)
	s.addMerge(ctx, 2, 4)
	s.addMerge(ctx, 1, 2)
	s.addMerge(ctx, 1, 3)
	s.addMerge(ctx, 2, 3)
	s.addMerge(ctx, 3, 15)
	return s.memDao.GetMemberByID(ctx, id)
}

func (s *Service) GetMemberByPhone(ctx context.Context, phone string) (member *model.Member, err error) {
	return s.memDao.GetMemberByPhone(ctx, phone)
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

func (s *Service) QueryMemberByName(ctx context.Context, name string) (members []*model.Member, err error) {
	return s.memDao.QueryMemberByName(ctx, name)
}

func (s *Service) QueryMemberByIDs(ctx context.Context, ids []int64) (res map[int64]*model.Member, err error) {
	return s.memDao.QueryMemberByIDs(ctx, ids)
}

func (s *Service) AddMember(ctx context.Context, member *model.Member) (id int64, err error) {
	id, err = s.memDao.AddMember(ctx, member)
	if err != nil {
		return
	}
	member.Id = id
	s.memDao.KafkaPushMember(ctx, member)
	return
}

func (s *Service) BatchAddMember(ctx context.Context, members []*model.Member) (affectRow int64, err error) {
	return s.memDao.BatchAddMember(ctx, members)
}

func (s *Service) InitMember(ctx context.Context, member *model.Member) (err error) {
	return s.memDao.InitMember(ctx, member)
}

func (s *Service) UpdateMember(ctx context.Context, member *model.Member) (err error) {
	return s.memDao.UpdateMember(ctx, member)
}

// attr demo
func (s *Service) UpdateMemberLock(ctx context.Context, id int64, lock int32) (err error) {
	mem, err := s.memDao.GetMemberByID(ctx, id)
	if err != nil {
		return
	}

	mem.AttrSet(lock, model.MemAttrLocked)
	return s.memDao.UpdateMemberAttr(ctx, id, mem.Attr)
}

func (s *Service) SetMember(ctx context.Context, member *model.Member) (err error) {
	return s.memDao.SetMember(ctx, member)
}

func (s *Service) SortMember(ctx context.Context, args model.ArgMemberSort) (err error) {
	return s.memDao.SortMember(ctx, args)
}

func (s *Service) DelMember(ctx context.Context, id int64) (err error) {
	return s.memDao.DeleteMember(ctx, id)
}

func (s *Service) addMerge(c context.Context, mid, kid int64) {
	s.merge.Add(c, strconv.FormatInt(mid, 10), &model.Merge{
		Mid: mid,
		Kid: kid,
	})
}

func (s *Service) initMerge() {
	s.merge = pipeline.NewPipeline(s.c.Merge)
	s.merge.Split = func(a string) int {
		n, _ := strconv.Atoi(a)
		return n
	}
	s.merge.Do = func(c context.Context, ch int, values map[string][]interface{}) {
		merges := make(map[string][]*model.Merge)
		for k, vs := range values {
			for _, v := range vs {
				merges[k] = append(merges[k], v.(*model.Merge))
			}
		}
		spew.Dump("merge start:", values)
		//log.Info("merges:%v" ,merges)
		// demo 数据库操作最终入库
		//s.dao.AddHistoryMessage(c, ch, merges)
	}
	s.merge.Start()
}
