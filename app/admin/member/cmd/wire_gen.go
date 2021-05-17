package main

import (
	"go-kartos-study/app/admin/member/conf"
	favDao "go-kartos-study/app/admin/member/internal/dao/favorite"
	memDao "go-kartos-study/app/admin/member/internal/dao/member"
	"go-kartos-study/app/admin/member/internal/server/http"
	"go-kartos-study/app/admin/member/internal/service/favorite"
	"go-kartos-study/app/admin/member/internal/service/member"
	"go-kartos-study/pkg/database/orm"
)

func initApp(c *conf.Config)(closeFunc func()) {
	db := orm.NewMySQL(c.ORM)
	db.LogMode(true)

	favDao := favDao.New(db)
	memDao := memDao.New(db)

	memService := member.New(memDao)
	favService := favorite.New(favDao)
	err := http.Init(c.HTTPServer,favService,memService)
	if err != nil{
		panic(err)
	}
	return  http.CloseService
}
