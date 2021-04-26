package member

import (
	"context"
	"go-kartos-study/app/bff/member/conf"
	"go-kartos-study/app/bff/member/internal/model"
	"go-kartos-study/app/service/member/api/grpc"
	"go-kartos-study/pkg/ecode"
	strx "go-kartos-study/pkg/str"
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

func (s *Service) GetMemberByID(ctx context.Context, arg *model.GetMemberByIDReq) (result *model.GetMemberResp, err error) {
	resp, err := s.memRPC.GetMemberByID(ctx, &grpc.GetMemberByIDReq{
		Id: arg.Id,
	})
	if err != nil {
		return nil, ecode.ErrQuery
	}
	result = &model.GetMemberResp{Member: model.Member{
		ID:        resp.Member.Id,
		Phone:     resp.Member.Phone,
		Name:      resp.Member.Name,
		Age:       resp.Member.Age,
		Address:   resp.Member.Address,
		CreatedAt: resp.Member.CreatedAt,
		UpdatedAt: resp.Member.UpdatedAt,
	}}
	return
}

func (s *Service) GetMemberByPhone(ctx context.Context, arg *model.GetMemberByPhoneReq) (result *model.GetMemberResp, err error) {
	resp, err := s.memRPC.GetMemberByPhone(ctx, &grpc.GetMemberByPhoneReq{
		Phone: arg.Phone,
	})
	if err != nil {
		return nil, ecode.ErrQuery
	}
	result = &model.GetMemberResp{Member: model.Member{
		ID:        resp.Member.Id,
		Phone:     resp.Member.Phone,
		Name:      resp.Member.Name,
		Age:       resp.Member.Age,
		Address:   resp.Member.Address,
		CreatedAt: resp.Member.CreatedAt,
		UpdatedAt: resp.Member.UpdatedAt,
	}}
	return
}

func (s *Service) GetMemberMaxAge(ctx context.Context) (result *model.GetMemberMaxAgeResq, err error) {
	resp, err := s.memRPC.GetMemberMaxAge(ctx, &grpc.GetMemberMaxAgeReq{
	})
	if err != nil {
		return nil, ecode.ErrQuery
	}
	result = &model.GetMemberMaxAgeResq{Age: resp.Age}
	return
}

func (s *Service) QueryMemberByName(ctx context.Context, req *model.QueryMemberByNameReq) (result *model.QueryMemberByNameResq, err error) {
	resp, err := s.memRPC.QueryMemberByName(ctx, &grpc.QueryMemberByNameReq{
		Name: req.Name,
	})
	if err != nil {
		return nil, ecode.ErrQuery
	}
	members := []model.Member{}
	for _, v := range resp.List {
		members = append(members, model.Member{
			ID:        v.Id,
			Phone:     v.Phone,
			Name:      v.Name,
			Age:       v.Age,
			Address:   v.Address,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}
	result = &model.QueryMemberByNameResq{
		Members: members,
	}
	return
}

func (s *Service) QueryMemberByIds(ctx context.Context, req *model.QueryMemberByIdsReq) (result *model.QueryMemberByIdsResp, err error) {
	ids, err := strx.SplitInts(req.IDs)
	if err != nil {
		return nil, ecode.ErrInvalidParam
	}
	resp, err := s.memRPC.QueryMemberByIDs(ctx, &grpc.QueryMemberByIDsReq{
		Ids: ids,
	})
	if err != nil {
		return nil, ecode.ErrQuery
	}
	list := make(map[int64]*model.Member)
	for k, v := range resp.List {
		list[k] = &model.Member{
			ID:        v.Id,
			Phone:     v.Phone,
			Name:      v.Name,
			Age:       v.Age,
			Address:   v.Address,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
	}
	result = &model.QueryMemberByIdsResp{
		Members: list,
	}
	return
}
func (s *Service) ListMember(ctx context.Context, req *model.ListMemberReq) (result *model.ListMemberResp, err error) {
	resp, err := s.memRPC.CountMember(ctx, &grpc.CountMemberReq{})
	if err != nil {
		return nil, ecode.ErrQuery
	}
	count := resp.Count
	resp2, err := s.memRPC.ListMember(ctx, &grpc.ListMemberReq{})
	if err != nil {
		return nil, ecode.ErrQuery
	}
	members := []model.Member{}
	for _, v := range resp2.List {
		members = append(members, model.Member{
			ID:        v.Id,
			Phone:     v.Phone,
			Name:      v.Name,
			Age:       v.Age,
			Address:   v.Address,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}
	result = &model.ListMemberResp{
		Page: &model.Page{
			Pn:    req.Pn,
			Ps:    req.Ps,
			Total: int(count),
		},
		Members: members,
	}
	return
}

func (s *Service) AddMember(ctx context.Context, req *model.AddMemberReq) (result *model.AddMemberResp, err error) {
	resp, err := s.memRPC.AddMember(ctx, &grpc.AddMemberReq{
		Phone:   req.Phone,
		Name:    req.Name,
		Age:     req.Age,
		Address: req.Address,
	})
	if err != nil {
		return nil, ecode.ErrInsert
	}
	result = &model.AddMemberResp{
		ID: resp.Id,
	}
	return
}

func (s *Service) InitMember(ctx context.Context, req *model.InitMemberReq) (err error) {
	_, err = s.memRPC.InitMember(ctx, &grpc.InitMemberReq{
		Phone:   req.Phone,
		Name:    req.Name,
		Age:     req.Age,
		Address: req.Address,
	})
	if err != nil {
		return ecode.ErrInsert
	}

	return
}

func (s *Service) BatchAddMember(ctx context.Context, req *model.BatchAddMemberReq) (err error) {
	var args []*grpc.AddMemberReq
	for _, v := range *req {
		args = append(args, &grpc.AddMemberReq{
			Phone:   v.Phone,
			Name:    v.Name,
			Age:     v.Age,
			Address: v.Address,
		})
	}

	_, err = s.memRPC.BatchAddMember(ctx, &grpc.BatchAddMemberReq{
		AddMemberReq: args,
	})
	if err != nil {
		return ecode.ErrInsert
	}
	return
}

func (s *Service) UpdateMember(ctx context.Context, req *model.UpdateMemberReq) (err error) {
	_, err = s.memRPC.UpdateMember(ctx, &grpc.UpdateMemberReq{
		Id:      req.ID,
		Phone:   req.Phone,
		Name:    req.Name,
		Age:     req.Age,
		Address: req.Address,
	})
	if err != nil {
		return ecode.ErrUpdate
	}
	return nil
}

func (s *Service) SetMember(ctx context.Context, req *model.SetMemberReq) (err error) {
	_, err = s.memRPC.SetMember(ctx, &grpc.SetMemberReq{
		Id:      req.ID,
		Phone:   req.Phone,
		Name:    req.Name,
		Age:     req.Age,
		Address: req.Address,
	})
	if err != nil {
		return ecode.ErrUpdate
	}
	return nil
}

func (s *Service) SortMember(ctx context.Context, req *model.SortMemberReq) (err error) {
	if len(req.Args) == 0 {
		return
	}
	args := make([]*grpc.SortMember, 0)
	for _, v := range req.Args {
		args = append(args, &grpc.SortMember{
			Id:       v.ID,
			OrderNum: v.OrderNum,
		})
	}
	_, err = s.memRPC.SortMember(ctx, &grpc.SortMemberReq{
		SortMember: args,
	})
	if err != nil {
		return ecode.ErrUpdate
	}
	return nil
}

func (s *Service) DelMember(ctx context.Context, req *model.DelMemberReq) error {
	_, err := s.memRPC.DeleteMember(ctx, &grpc.DeleteMemberReq{
		Id: req.ID,
	})
	if err != nil {
		return ecode.ErrDelete
	}
	return nil
}
