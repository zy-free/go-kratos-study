package model

import xtime "go-kartos-study/pkg/time"

type Favorite struct {
	ID     int64      `json:"id"`
	Mid    int64      `json:"mid"`  // 用户id
	Name   string     `json:"name"` // 收藏夹名称
	HintAt xtime.Time `json:"hint_at"`
}
