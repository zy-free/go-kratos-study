package jaeger

import (
	"flag"
	"go-kartos-study/pkg/conf/env"
	"os"

	"go-kartos-study/pkg/net/trace"
)

var (
	_jaegerAppID    = env.AppID
	//_jaegerAppID    = "bff-member-service"
	_jaegerEndpoint = "http://*:14268/api/traces"
)

func init() {
	if v := os.Getenv("JAEGER_ENDPOINT"); v != "" {
		_jaegerEndpoint = v
	}

	if v := os.Getenv("JAEGER_APPID"); v != "" {
		_jaegerAppID = v
	}

	flag.StringVar(&_jaegerEndpoint, "jaeger_endpoint", _jaegerEndpoint, "jaeger report endpoint, or use JAEGER_ENDPOINT env.")
	flag.StringVar(&_jaegerAppID, "jaeger_appid", _jaegerAppID, "jaeger report appid, or use JAEGER_APPID env.")
}

// Init Init
func Init() {
	c := &Config{Endpoint: _jaegerEndpoint, BatchSize: 120}
	trace.SetGlobalTracer(trace.NewTracer(_jaegerAppID, newReport(c), true))
}
