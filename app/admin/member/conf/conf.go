package conf

import (
	"flag"
	"go-kartos-study/pkg/database/orm"
	"go-kartos-study/pkg/log"
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

	ORM    *orm.Config
	HTTPClient *bm.ClientConfig

}

func init(){
	flag.StringVar(&confPath, "conf", "./app/admin/member/cmd/config.toml", "default config path")
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
