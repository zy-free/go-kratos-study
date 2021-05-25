package grpc

import (
	"context"
	"fmt"
	"go-kartos-study/app/service/member/appid"
	"go-kartos-study/pkg/net/rpc/warden"

	"google.golang.org/grpc"
)

// NewClient include TradeClient and SalesClient
func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (MemberRPCClient, error) {
	// dis, err := New(&clientv3.Config{Endpoints: conf.Endpoints})
	// if err != nil {
	//	panic(err)
	// }
	cc, err := warden.NewClient(cfg, opts...).Dial(context.Background(), fmt.Sprintf("etcd://default/%s", appid.AppID))
	if err != nil {
		return nil, err
	}
	return NewMemberRPCClient(cc), nil
}
