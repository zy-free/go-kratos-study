package grpc

import (
	"context"
	"go-kartos-study/app/service/member/api/grpc"
	"go-kartos-study/app/service/member/internal/model"
)

func (s *Server) GetMemberByID(ctx context.Context, req *grpc.GetMemberByIDReq) (resp *grpc.MemberResp, err error) {
	member,err := s.svc.GetMemberByID(ctx,req.Id)
	if err != nil {
		return nil, err
	}

	resp = &grpc.MemberResp{
		Member:&grpc.Member{
			Id:        member.Id,
			Phone:     member.Phone,
			Name:      member.Name,
			Age:       member.Age,
			Address:   member.Address,
			CreatedAt: member.CreatedAt.Unix(),
			UpdatedAt: member.UpdatedAt.Unix(),
		},
	}
	return
}

func (s *Server) GetMemberByPhone(ctx context.Context, req *grpc.GetMemberByPhoneReq) (resp *grpc.MemberResp, err error) {
	member,err := s.svc.GetMemberByPhone(ctx,req.Phone)
	if err != nil {
		return nil, err
	}

	resp = &grpc.MemberResp{
		Member:&grpc.Member{
			Id:        member.Id,
			Phone:     member.Phone,
			Name:      member.Name,
			Age:       member.Age,
			Address:   member.Address,
			CreatedAt: member.CreatedAt.Unix(),
			UpdatedAt: member.UpdatedAt.Unix(),
		},
	}
	return
}

func (s *Server) GetMemberMaxAge(ctx context.Context, req *grpc.GetMemberMaxAgeReq) (resp *grpc.GetMemberMaxAgeResp, err error) {
	age, err := s.svc.GetMemberMaxAge(ctx)
	if err != nil {
		return nil, err
	}

	return &grpc.GetMemberMaxAgeResp{
		Age: age,
	}, nil
}
func (s *Server) CountMember(ctx context.Context, req *grpc.CountMemberReq) (resp *grpc.CountMemberResp, err error) {
	count, err := s.svc.CountMember(ctx)
	if err != nil {
		return nil, err
	}

	return &grpc.CountMemberResp{
		Count: count,
	}, nil
}
func (s *Server) ListMember(ctx context.Context, req *grpc.ListMemberReq) (resp *grpc.ListMemberResp, err error) {
	members, err := s.svc.ListMember(ctx)
	if err != nil {
		return nil, err
	}
	var mResp []*grpc.Member
	for _, member := range members {
		mResp = append(mResp, &grpc.Member{
			Id:        member.Id,
			Phone:     member.Phone,
			Name:      member.Name,
			Age:       member.Age,
			Address:   member.Address,
			CreatedAt: member.CreatedAt.Unix(),
			UpdatedAt: member.UpdatedAt.Unix(),
		})
	}
	return &grpc.ListMemberResp{
		List: mResp,
	}, nil
}

func (s *Server) QueryMemberByName(ctx context.Context, req *grpc.QueryMemberByNameReq) (resp *grpc.QueryMemberByNameResp, err error) {
	members, err := s.svc.QueryMemberByName(ctx,req.Name)
	if err != nil {
		return nil, err
	}

	var mResp []*grpc.Member
	for _, member := range members {
		mResp = append(mResp, &grpc.Member{
			Id:        member.Id,
			Phone:     member.Phone,
			Name:      member.Name,
			Age:       member.Age,
			Address:   member.Address,
			CreatedAt: member.CreatedAt.Unix(),
			UpdatedAt: member.UpdatedAt.Unix(),
		})
	}
	return &grpc.QueryMemberByNameResp{
		List: mResp,
	}, nil
}

func (s *Server) QueryMemberByIDs(ctx context.Context, req *grpc.QueryMemberByIDsReq) (resp *grpc.QueryMemberByIDsResp, err error) {
	members, err := s.svc.QueryMemberByIDs(ctx,req.Ids)
	if err != nil {
		return nil, err
	}

	mResp := make(map[int64]*grpc.Member)
	for k, member := range members {
		mResp[k] = &grpc.Member{
			Id:        member.Id,
			Phone:     member.Phone,
			Name:      member.Name,
			Age:       member.Age,
			Address:   member.Address,
			CreatedAt: member.CreatedAt.Unix(),
			UpdatedAt: member.UpdatedAt.Unix(),
		}
	}
	return &grpc.QueryMemberByIDsResp{
		List: mResp,
	}, nil
}
func (s *Server) AddMember(ctx context.Context, req *grpc.AddMemberReq) (resp *grpc.IDResp, err error) {
	m := &model.Member{
		Phone:   req.Phone,
		Name:    req.Name,
		Age:     req.Age,
		Address: req.Address,
	}
	// 业务端不关心，一定做好转化 分割字段，Get返回字段时同理
	//m.AttrSet(model.MemAttrPublic,req.Public)
	//m.AttrSet(model.MemAttrLocked,req.Locked)
	id, err := s.svc.AddMember(ctx,m)
	if err != nil {
		return nil, err
	}
	return &grpc.IDResp{
		Id: id,
	}, nil
}
func (s *Server) BatchAddMember(ctx context.Context, req *grpc.BatchAddMemberReq) (resp *grpc.BatchAddMemberResp, err error) {
	var args []*model.Member
	for _, v := range req.AddMemberReq {
		args = append(args, &model.Member{
			Phone:   v.Phone,
			Name:    v.Name,
			Age:     v.Age,
			Address: v.Address,
		})
	}
	affectRow, err := s.svc.BatchAddMember(ctx,args)
	if err != nil {
		return nil, err
	}
	return &grpc.BatchAddMemberResp{
		AffectRow: affectRow,
	}, nil
}

func (s *Server) InitMember(ctx context.Context, req *grpc.InitMemberReq) (resp *grpc.EmptyResp, err error) {
	err = s.svc.InitMember(ctx,&model.Member{
		Phone:   req.Phone,
		Name:    req.Name,
		Age:     req.Age,
		Address: req.Address,
	})
	if err != nil {
		return nil, err
	}
	return &grpc.EmptyResp{}, nil
}

func (s *Server) DeleteMember(ctx context.Context, req *grpc.DeleteMemberReq) (resp *grpc.EmptyResp, err error) {
	err = s.svc.DelMember(ctx,req.Id)
	if err != nil {
		return nil, err
	}
	return &grpc.EmptyResp{}, nil
}

func (s *Server) SetMember(ctx context.Context, req *grpc.SetMemberReq) (resp *grpc.EmptyResp, err error) {
	err = s.svc.SetMember(ctx,&model.Member{
		Id:      req.Id,
		Phone:   req.Phone,
		Name:    req.Name,
		Age:     req.Age,
		Address: req.Address,
	})
	return &grpc.EmptyResp{}, err
}

func (s *Server) UpdateMember(ctx context.Context, req *grpc.UpdateMemberReq) (resp *grpc.EmptyResp, err error) {
	err = s.svc.UpdateMember(ctx,&model.Member{
		Id:      req.Id,
		Phone:   req.Phone,
		Name:    req.Name,
		Age:     req.Age,
		Address: req.Address,
	})
	return &grpc.EmptyResp{}, err
}

func (s *Server) SortMember(ctx context.Context, in *grpc.SortMemberReq) (resp *grpc.EmptyResp, err error) {
	var args model.ArgMemberSort
	for _, v := range in.SortMember {
		args = append(args,model.MemberSort{Id: v.Id, OrderNum: v.OrderNum})
	}
	err = s.svc.SortMember(ctx,args)
	return &grpc.EmptyResp{}, err
}
