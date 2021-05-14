package model

import "go-kartos-study/pkg/hashid"

type Page struct {
	Pn    int `json:"pn"`
	Ps    int `json:"ps"`
	Total int `json:"total"`
}

type GetMemberByIDReq struct {
	Id int64 `form:"id" validate:"required"`
}

type GetMemberByPhoneReq struct {
	Phone string `form:"phone"`
}
type GetMemberResp struct {
	Member
}
type GetMemberMaxAgeResq struct {
	Age int64 `json:"age"`
}
type QueryMemberByNameReq struct {
	Name string `form:"name"`
}
type QueryMemberByNameResq struct {
	Members []Member `json:"members"`
}
type QueryMemberByIdsReq struct {
	IDs string `form:"ids"`
}
type QueryMemberByIdsResp struct {
	Members map[hashid.ID]*Member `json:"members"`
}

type ListMemberReq struct {
	Pn int `form:"pn" default:"1"`
	Ps int `form:"ps" default:"15"`
}

type ListMemberResp struct {
	Page    *Page    `json:"page"`
	Members []Member `json:"members"`
}

type AddMemberReq struct {
	ID      int64  `json:"id"`
	Phone   string `json:"phone" validate:"required"`
	Name    string `json:"name"`
	Age     int64  `json:"age"`
	Address string `json:"address"`
}

type AddMemberResp struct {
	ID int64 `json:"id"`
}

type InitMemberResp struct {
	ID int64 `json:"id"`
}

type BatchAddMemberReq struct {
	Args []*AddMemberReq
}

type UpdateMemberReq struct {
	ID      int64  `json:"id"  validate:"required" default:"3"`
	Phone   string `json:"phone"  `
	Name    string `json:"name"  default:"-1"`
	Age     int64  `json:"age"   default:"-1"`
	Address string `json:"address" default:"-1"`
}

type SetMemberReq struct {
	ID      int64  `json:"id" validate:"required"`
	Phone   string `json:"phone"`
	Name    string `json:"name"`
	Age     int64  `json:"age"`
	Address string `json:"address"`
}

type SortMember struct {
	ID       int64 `json:"id"`
	OrderNum int64 `json:"order_num"`
}

type SortMemberReq struct {
	Args []SortMember `json:"args"`
}

type DelMemberReq struct {
	ID int64 `json:"id" validate:"required"`
}

type GetFavoriteByIDReq struct {
	Id int64 `form:"id" validate:"required"`
}
type GetFavoriteResp struct {
	Favorite
}
type AddFavoriteReq struct {
	Name string `json:"name"`
	Mid  int64  `json:"mid"`
}

type AddFavoriteResp struct {
	ID int64 `json:"id"`
}

type MemberSort struct {
	Id       int64 `db:"id"`
	OrderNum int64 `db:"order_num"`
}
type ArgMemberSort []MemberSort
