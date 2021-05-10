package conf

import (
	"flag"
	"go-kartos-study/pkg/log"
	"go-kartos-study/pkg/net/rpc/warden"
	"go-kartos-study/pkg/net/trace"

	"github.com/BurntSushi/toml"
	bm "go-kartos-study/pkg/net/http/blademaster"
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

// Init init conf
func Init() error {
	flag.StringVar(&confPath, "conf", "./app/bff/member/cmd/config.toml", "default config path")
	if confPath != "" {
		return local()
	}
	return nil
}

func local() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}
