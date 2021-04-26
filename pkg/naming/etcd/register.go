package etcd

import (
	"context"
	"github.com/google/uuid"
	"go-kartos-study/pkg/naming"
	"go.etcd.io/etcd/clientv3"
	"os"
)

type ETCDConfig struct {
	Endpoints []string
}

func ETCDRegist(conf *ETCDConfig, appID string, port string, color string) (cancel context.CancelFunc) {
	dis, err := New(&clientv3.Config{Endpoints: conf.Endpoints})
	if err != nil {
		panic(err)
	}

	hn, _ := os.Hostname()

	ins := &naming.Instance{
		AppID: appID,
		Addrs: []string{
			//"grpc://" + ip.InternalIP() + ":" + port,
			"grpc://" + "127.0.0.1" + ":" + port,
		},
		Metadata: map[string]string{"color": color},
		Hostname: hn + uuid.New().String(),
	}
	cancel, err = dis.Register(context.Background(), ins)
	if err != nil {
		panic(err)
	}
	//defer cancel()
	return
}
