package model

import "time"

type Member struct {
	Id        int64     `db:"id"`
	Phone     string    `db:"phone"`
	Name      string    `db:"name"`
	Age       int64     `db:"age"`
	Address   string    `db:"address"`
	OrderNum  int64     `db:"order_num"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
type MemberSort struct {
	Id       int64 `db:"id"`
	OrderNum int64 `db:"order_num"`
}
type ArgMemberSort []MemberSort
