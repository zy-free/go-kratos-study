package conf

import (
	"flag"
	"go-kartos-study/pkg/log"
	"go-kartos-study/pkg/net/trace"
	"go-kartos-study/pkg/queue/kafka"

	"github.com/BurntSushi/toml"
)

// global var
var (
	confPath string
	// Conf config
	Conf = &Config{}
)

// Config config set
type Config struct {
	Log           *log.Config
	Tracer        *trace.Config
	KafkaConsumer *kafka.ConsumerConfig
}

// Init init conf
func Init() error {
	flag.StringVar(&confPath, "conf", "./app/job/member/cmd/config.toml", "default config path")
	if confPath != "" {
		return local()
	}
	return nil
}

func local() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}
