package model

type Favorite struct {
	Id        int64     `db:"id"`
	Mid       int64     `db:"mid"`        // 用户id
	Name      string    `db:"name"`       // 收藏夹名称
}
