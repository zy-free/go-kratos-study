package model

import (
	"go-kartos-study/pkg/hashid"
	"strconv"
)

type Member struct {
	ID        hashid.ID `json:"id"`
	Phone     string    `json:"phone"`
	Name      string    `json:"name"`
	Age       int64     `json:"age"`
	Address   string    `json:"address"`
	OrderNum  int64     `json:"order_num"`
	CreatedAt int64     `json:"-"`
	UpdatedAt int64     `json:"-"`
}

// 贫血模型

// 判断用户是否old
func (v *Member) IsOld() bool {
	return v.Age >= 35
}

func (v *Member) IsOldToString() string {
	var isOld = "否"
	if v.IsOld() {
		isOld = "是"
	}
	return isOld
}

func MemberExportFields() []string {
	return []string{"ID", "手机号", "名字", "年龄", "是否老龄", "地址"}
}

func (v *Member) ExportStrings() []string {
	return []string{
		strconv.FormatInt(int64(v.ID), 10),
		v.Phone,
		v.Name,
		"年龄：" + strconv.FormatInt(v.Age, 10),
		v.IsOldToString(),
		v.Address,
	}
}
