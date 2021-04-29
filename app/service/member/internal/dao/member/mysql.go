package member

import (
	"bytes"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go-kartos-study/app/service/member/internal/model"
	"go-kartos-study/pkg/database/sql"
	xstr "go-kartos-study/pkg/str"
	"strconv"
	"strings"
	"time"
)

const (
	_shard = 100
)

// 分表命名:表名+hit
func memberHit(id int64) string {
	return fmt.Sprintf("member_%d", id%_shard)
}

func (dao *Dao) dbInitMember(ctx context.Context, arg *model.Member) (err error) {
	_sql := `insert ignore into member (phone,name,age,address) VALUES(?,?,?,?)`
	if _, err = dao.db.Exec(ctx, _sql, arg.Phone, arg.Name, arg.Age, arg.Address); err != nil {
		return errors.Wrapf(err, "InitMember arg(%v)", arg)
	}
	return
}

// 创建单个
func (dao *Dao) dbAddMember(ctx context.Context, arg *model.Member) (id int64, err error) {
	_sql := `INSERT INTO member (phone,name,age,address) VALUES(?,?,?,?)`
	result, err := dao.db.Exec(ctx, _sql, arg.Phone, arg.Name, arg.Age, arg.Address)
	if err != nil {
		return 0, errors.Wrapf(err, "AddMember arg(%v)", arg)
	}

	id, _ = result.LastInsertId()
	return
}

// 批量创建
func (dao *Dao) dbBatchAddMember(ctx context.Context, args []*model.Member) (affectRow int64, err error) {
	_sql := `INSERT INTO member (phone,name,age,address) VALUES `
	var valueString []string
	var valueArgs []interface{}
	for _, arg := range args {
		valueString = append(valueString, "(?,?,?,?)")
		valueArgs = append(valueArgs, arg.Phone, arg.Name, arg.Age, arg.Address)
	}
	result, err := dao.db.Exec(ctx, _sql+strings.Join(valueString, ","), valueArgs...)
	if err != nil {
		return 0, errors.Wrapf(err, "BatchAddMember arg(%v)", args)
	}
	affectRow, _ = result.RowsAffected()
	return
}

// 根据id查询单个
func (dao *Dao) dbGetMemberByID(ctx context.Context, id int64) (m *model.Member, err error) {
	m = &model.Member{}
	_sql := `SELECT id,phone,name,age,address FROM member WHERE id = ? AND deleted_at is null `
	if err = dao.db.QueryRow(ctx, _sql, id).Scan(&m.Id, &m.Phone, &m.Name, &m.Age, &m.Address); err != nil {
		return nil, errors.Wrapf(err, "GetMemberByID id(%d)", id)
	}
	return
}

// 根据phone查询单个
func (dao *Dao) dbGetMemberByPhone(ctx context.Context, phone string) (m *model.Member, err error) {
	m = &model.Member{}
	_sql := `SELECT id,phone,name,age,address FROM member WHERE phone = ? AND deleted_at is null `
	if err = dao.db.QueryRow(ctx, _sql, phone).Scan(&m.Id, &m.Phone, &m.Name, &m.Age, &m.Address); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			return nil, errors.Wrapf(err, "GetMemberByPhone phone(%s)", phone)
		}
	}
	return
}


func (dao *Dao) dbGetMemberMaxAge(ctx context.Context) (age int64, err error) {
	_sql := `SELECT IFNULL(MAX(age),0) FROM member WHERE deleted_at is null `
	if err = dao.db.QueryRow(ctx, _sql).Scan(&age); err != nil {
		return 0, errors.Wrapf(err, "GetMemberMaxAge")
	}
	return
}

func (dao *Dao) dbGetMemberSumAge(ctx context.Context) (age int64, err error) {
	_sql := `SELECT IFNULL(SUM(age),0) FROM member WHERE  deleted_at is null  `
	if err = dao.db.QueryRow(ctx, _sql).Scan(&age); err != nil {
		return 0, errors.Wrapf(err, "GetMemberMaxAge")
	}
	return
}

func (dao *Dao) dbCountMember(ctx context.Context) (count int64, err error) {
	_sql := `SELECT COUNT(*) FROM member where  deleted_at is null `
	if err = dao.db.QueryRow(ctx, _sql).Scan(&count); err != nil {
		return 0, errors.Wrapf(err, "CountMember")
	}
	return
}

func (dao *Dao) dbListMember(ctx context.Context) (res []*model.Member, err error) {
	res = make([]*model.Member, 0, 0) // 返回nil还是空切片会影响json里的结构
	_sql := "SELECT id,phone,name,age,address FROM member WHERE  deleted_at is null "
	rows, err := dao.db.Query(ctx, _sql,)
	if err != nil {
		return res, errors.Wrapf(err, "ListMember ")
	}
	for rows.Next() {
		n := &model.Member{}
		if err = rows.Scan(&n.Id, &n.Phone, &n.Name, &n.Age, &n.Address); err != nil {
			err = errors.Wrapf(err, "d.db.Scan(%s)")
			return
		}
		res = append(res, n)
	}
	if err = rows.Err(); err != nil {
		err = errors.Wrapf(err, "rows.Err(%s)")
	}
	return
}

func (dao *Dao) dbHasMemberByID(ctx context.Context, id int64) (has bool, err error) {
	var count int
	_sql := `SELECT COUNT(*) FROM member WHERE id = ?  AND deleted_at is null  `
	err = dao.db.QueryRow(ctx, _sql, id).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}
		return false, errors.Wrapf(err, "HasMemberByID id(%d)", id)
	}
	return count > 0, nil
}

// 根据其他属性查询列表
func (dao *Dao) dbQueryMemberByName(ctx context.Context, name string) (res []*model.Member, err error) {
	res = make([]*model.Member, 0, 0) // 返回nil还是空切片会影响json里的结构
	_sql := "SELECT id,phone,name,age,address FROM member WHERE name = ? AND deleted_at is null "
	rows, err := dao.db.Query(ctx, _sql, name)
	if err != nil {
		return res, errors.Wrapf(err, "QueryMemberByName name(%s)", name)
	}
	for rows.Next() {
		n := &model.Member{}
		if err = rows.Scan(&n.Id, &n.Phone, &n.Name, &n.Age, &n.Address); err != nil {
			err = errors.Wrapf(err, "d.db.Scan(%s)", name)
			return
		}
		res = append(res, n)
	}
	if err = rows.Err(); err != nil {
		err = errors.Wrapf(err, "rows.Err(%s)", name)
	}
	return
}

// 根据ids查询列表
func (dao *Dao) dbQueryMemberByIDs(ctx context.Context, ids []int64) (res map[int64]*model.Member, err error) {
	res = make(map[int64]*model.Member)
	_sql := "SELECT id,phone,name,age,address FROM member WHERE id IN (" + xstr.JoinInts(ids) + ") AND deleted_at is null "
	rows, err := dao.db.Query(ctx, _sql)
	if err != nil {
		return res, errors.Wrapf(err, "QueryMemberByIDs name(%s)", ids)
	}
	for rows.Next() {
		n := &model.Member{}
		if err = rows.Scan(&n.Id, &n.Phone, &n.Name, &n.Age, &n.Address); err != nil {
			err = errors.Wrapf(err, "d.db.Scan(%s)", ids)
			return
		}
		res[n.Id] = n
	}
	if err = rows.Err(); err != nil {
		err = errors.Wrapf(err, "rows.Err(%s)", ids)
	}
	return
}


// 更新单个
func (dao *Dao) dbUpdateMember(ctx context.Context, member *model.Member) (err error) {
	_sql := "UPDATE member SET  "
	sqlSli := []string{}
	var updateMap []interface{}

	if member.Phone != "-1" {
		sqlSli = append(sqlSli, "phone =?")
		updateMap = append(updateMap, member.Phone)
	}
	if member.Name != "-1" {
		sqlSli = append(sqlSli, "name =?")
		updateMap = append(updateMap, member.Name)
	}
	if member.Address != "-1" {
		sqlSli = append(sqlSli, "address =?")
		updateMap = append(updateMap, member.Address)
	}
	if member.Age != -1 {
		sqlSli = append(sqlSli, "age =?")
		updateMap = append(updateMap, member.Age)
	}
	_sql += strings.Join(sqlSli, ",") + " WHERE id =?"
	updateMap = append(updateMap, member.Id)
	if _, err = dao.db.Exec(ctx, _sql, updateMap...); err != nil {
		return errors.Wrapf(err, "Update arg(%v)", updateMap)
	}
	return
}

// 更新或创建
func (dao *Dao) dbSetMember(ctx context.Context, arg *model.Member) (err error) {
	_sql := "INSERT INTO member (id,phone,name,age,address) VALUES (?,?,?,?,?) " +
		"ON DUPLICATE KEY UPDATE phone=?,name=?,age=?,address=?"
	if _, err = dao.db.Exec(ctx, _sql, arg.Id, arg.Phone, arg.Name, arg.Age, arg.Address, arg.Phone, arg.Name, arg.Age, arg.Address); err != nil {
		return errors.Wrapf(err, "SetMember arg(%v)", arg)
	}
	return
}

// 批量更改顺序
func (dao *Dao) dbSortMember(ctx context.Context, args model.ArgMemberSort) (err error) {
	var (
		buf bytes.Buffer
		ids []int64
	)
	buf.WriteString("UPDATE member SET order_num = CASE id")
	for _, arg := range args {
		buf.WriteString(" WHEN ")
		buf.WriteString(strconv.FormatInt(arg.Id, 10))
		buf.WriteString(" THEN ")
		buf.WriteString(strconv.FormatInt(arg.OrderNum, 10))
		ids = append(ids, arg.Id)
	}
	buf.WriteString(" END  WHERE id IN (")
	buf.WriteString(xstr.JoinInts(ids))
	buf.WriteString(")")
	if _, err = dao.db.Exec(ctx, buf.String()); err != nil {
		return errors.Wrapf(err, "BatchUpdateMemberOrder args(%v)", args)
	}
	return
}


// 删除
func (dao *Dao) dbDeleteMember(ctx context.Context, id int64) (err error) {
	now := time.Now()
	_sql := `UPDATE member SET deleted_at = ? WHERE id = ?`
	if _, err = dao.db.Exec(ctx, _sql, now, id); err != nil {
		return errors.Wrapf(err, "DeleteMember id(%d)", id)
	}
	return
}
