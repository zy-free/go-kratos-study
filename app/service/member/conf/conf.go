package conf

import (
	"flag"
	"github.com/BurntSushi/toml"
	"go-kartos-study/pkg/cache/redis"
	"go-kartos-study/pkg/database/sql"
	"go-kartos-study/pkg/log"
	"go-kartos-study/pkg/naming/etcd"
	"go-kartos-study/pkg/net/rpc/warden"
	"go-kartos-study/pkg/net/trace"
	kafka "go-kartos-study/pkg/queue/kafka"
	"go-kartos-study/pkg/sync/pipeline"
)

// global var
var (
	confPath string
	// Conf config
	Conf = &Config{}
)

// Config config set
type Config struct {
	Log    *log.Config
	Tracer *trace.Config
	Mysql  *sql.Config
	Redis  *redis.Config
	Merge  *pipeline.Config
	KafkaPublish  *kafka.PublishConfig

	GRPCServer *warden.ServerConfig
	ETCDConfig *etcd.ETCDConfig
}

func init() {
	flag.StringVar(&confPath, "conf", "./app/service/member/cmd/config.toml", "default config path")
}

// Init init conf
func Init() error {
	if confPath != "" {
		return local()
	}
	return nil
}

func local() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}
