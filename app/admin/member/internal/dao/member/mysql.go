package member

import (
	"bytes"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go-kartos-study/app/admin/member/internal/model"
	"go-kartos-study/pkg/ecode"
)

const (
	_shard = 100
)

// 分表命名:表名+hit
func memberHit(id int64) string {
	return fmt.Sprintf("member_%d", id%_shard)
}

// 创建单个
func (dao *Dao) dbAddMember(ctx context.Context, arg *model.AddMemberReq) (id int64, err error) {
	err = dao.db.Table("member").Create(arg).Error
	if err != nil {
		return
	}
	id = arg.ID
	return
}

// 批量创建
func (dao *Dao) dbBatchAddMember(ctx context.Context, args []*model.AddMemberReq) (affectRow int64, err error) {
	var (
		buf  bytes.Buffer
		_sql string
	)
	buf.WriteString(`INSERT INTO member (phone,name,age,address) VALUES `)
	for _, v := range args {
		buf.WriteString("('")
		buf.WriteString(v.Phone)
		buf.WriteString("','")
		buf.WriteString(v.Name)
		buf.WriteString("','")
		buf.WriteString(fmt.Sprintf("%d", v.Age))
		buf.WriteString("','")
		buf.WriteString(v.Address)
		buf.WriteString("'),")
	}
	_sql = buf.String()
	err = dao.db.Exec(_sql[0 : len(_sql)-1]).Error
	if err != nil {
		return 0, errors.WithMessagef(err, "BatchAddMember arg(%v)", args)
	}
	return
}

// 根据id查询单个
func (dao *Dao) dbGetMemberByID(ctx context.Context, id int64) (m *model.Member, err error) {
	m = new(model.Member)
	if err = dao.db.Select("id,phone,name,age,address,attr").Table("member").Where("id=?", id).First(m).Error; err != nil {
		if err == ecode.NothingFound {
			err = nil
			return
		}
		err = errors.Wrapf(err, "dbGetMemberByID")
	}
	return
}

// 更新单个
func (dao *Dao) dbUpdateMember(ctx context.Context, member *model.UpdateMemberReq) (err error) {
	m := make(map[string]interface{})
	if member.Phone != "-1" {
		m["phone"] = member.Phone
	}
	if member.Name != "-1" {
		m["name"] = member.Name
	}
	if member.Address != "-1" {
		m["address"] = member.Address
	}
	if member.Age != -1 {
		m["age2"] = member.Age
	}
	err = dao.db.Table("member").Where("id=?", member.ID).Update(m).Error

	return err
}
