package hbase

import (
	xtime "go-kartos-study/pkg/time"
)

// ZKConfig Server&Client settings.
type ZKConfig struct {
	Root    string
	Addrs   []string
	Timeout xtime.Duration
}

// Config hbase config
type Config struct {
	Zookeeper           *ZKConfig
	RPCQueueSize        int
	FlushInterval       xtime.Duration
	EffectiveUser       string
	RegionLookupTimeout xtime.Duration
	RegionReadTimeout   xtime.Duration
	TestRowKey          string
}
