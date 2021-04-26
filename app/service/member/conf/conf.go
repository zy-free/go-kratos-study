package conf

import (
	"flag"
	"go-kartos-study/pkg/database/sql"
	"go-kartos-study/pkg/log"
	"go-kartos-study/pkg/naming/etcd"
	"go-kartos-study/pkg/net/rpc/warden"
	"go-kartos-study/pkg/net/trace"

	"github.com/BurntSushi/toml"
)

const(
	AppID = "member.service"
	Color = ""
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
	Mysql *sql.Config
	GRPCServer *warden.ServerConfig
	ETCDConfig *etcd.ETCDConfig
}


// Init init conf
func Init() error {
	flag.StringVar(&confPath, "conf", "./app/service/member/cmd/config.toml", "default config path")
	if confPath != "" {
		return local()
	}
	return nil
}

func local() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}
