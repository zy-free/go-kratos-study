package grpc

import (
	"context"

	"github.com/pkg/errors"

	"go-kartos-study/app/service/member/api/grpc"
	"go-kartos-study/pkg/ecode"
	"go-kartos-study/pkg/log"
	"go-kartos-study/pkg/net/metadata"
)

func (s *Server) ErrorTest(c context.Context, in *grpc.EmptyReq) (resp *grpc.EmptyResp, err error) {
	// client can get ecode.ErrConfigureReward,also can if err == ecode.ErrConfigureReward
	return nil, errors.Wrapf(ecode.ErrConfigureReward, "ErrorTest")
}

func (s *Server) MetadataTest(c context.Context, in *grpc.EmptyReq) (resp *grpc.EmptyResp, err error) {
	log.Info("mid:%s-----color:%s", metadata.String(c, metadata.Mid), metadata.String(c, metadata.Color))
	return &grpc.EmptyResp{}, nil
}
