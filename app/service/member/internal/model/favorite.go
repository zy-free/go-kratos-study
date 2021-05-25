package model

import xtime "go-kartos-study/pkg/time"

type Favorite struct {
	ID     int64      `db:"id"`
	Mid    int64      `db:"mid"`  // 用户id
	Name   string     `db:"name"` // 收藏夹名称
	HintAt xtime.Time `db:"hint_at"`
}
