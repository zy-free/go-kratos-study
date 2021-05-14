package model

import (
	"time"
)

const (
	// ResAttrLocked resource attr of locked
	MemAttrLocked = uint(0)
	MemAttrPublic = uint(1)

	AttrLockedNo  = int32(0)
	AttrLockedYes = int32(1)
	AttrPublic    = int32(0)
)

type Member struct {
	Id        int64     `db:"id"`
	Phone     string    `db:"phone"`
	Name      string    `db:"name"`
	Age       int64     `db:"age"`
	Address   string    `db:"address"`
	OrderNum  int64     `db:"order_num"`
	Attr      int32     `db:"attr"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// AttrVal gets attr val by bit.
func (r *Member) AttrVal(bit uint) int32 {
	return (r.Attr >> bit) & int32(1)
}

// AttrSet sets attr value by bit.
func (r *Member) AttrSet(v int32, bit uint) {
	r.Attr = r.Attr&(^(1 << bit)) | (v << bit)
}

// Locked resource locked state
func (r *Member) Locked() bool {
	return r.AttrVal(MemAttrLocked) == AttrLockedYes
}

func (r *Member) IsPublic() bool {
	return r.AttrVal(MemAttrPublic) == AttrPublic
}

type MemberSort struct {
	Id       int64 `db:"id"`
	OrderNum int64 `db:"order_num"`
}
type ArgMemberSort []MemberSort
