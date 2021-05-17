package main

import (
	"go-kartos-study/app/service/member/conf"
	favoriteDao "go-kartos-study/app/service/member/internal/dao/favorite"
	memberDao "go-kartos-study/app/service/member/internal/dao/member"
	"go-kartos-study/app/service/member/internal/server/grpc"
	"go-kartos-study/app/service/member/internal/service"
	"go-kartos-study/pkg/cache/redis"
	"go-kartos-study/pkg/database/sql"
	"go-kartos-study/pkg/queue/kafka"
	"go-kartos-study/pkg/sync/pipeline"
)

func initApp(c *conf.Config) (closeFunc func()){
	db := sql.NewMySQL(c.Mysql)
	redis := redis.NewPool(c.Redis)
	publisher := kafka.NewPublisher(c.KafkaPublish)
	merge := pipeline.NewPipeline(c.Merge)

	favDao := favoriteDao.New(db)
	memDao := memberDao.New(db, redis, publisher)
	svc := service.New(favDao, memDao, merge)

	grpc.New(conf.Conf.GRPCServer, svc)
	return svc.Close
}
