package conf

import (
	"flag"

	"github.com/BurntSushi/toml"

	"go-kartos-study/pkg/log"
	bm "go-kartos-study/pkg/net/http/blademaster"
	"go-kartos-study/pkg/net/rpc/warden"
	"go-kartos-study/pkg/net/trace"
)

// global var
var (
	confPath string
	// Conf config
	Conf = &Config{}
)

// Config config set
type Config struct {
	HTTPServer *bm.ServerConfig
	Log        *log.Config
	Tracer     *trace.Config

	MemberClient *warden.ClientConfig

	HTTPClient *bm.ClientConfig
}

func init() {
	flag.StringVar(&confPath, "conf", "./app/bff/member/cmd/config.toml", "default config path")
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
