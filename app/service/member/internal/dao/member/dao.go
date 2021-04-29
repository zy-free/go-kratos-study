package member

import (
	"context"
	"go-kartos-study/app/service/member/conf"
	"go-kartos-study/app/service/member/internal/model"
	"go-kartos-study/pkg/database/sql"
)

// Dao is redis dao.
type Dao struct {
	c *conf.Config
	// db
	db *sql.DB
}

// New new a dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:  c,
		db: sql.NewMySQL(c.Mysql),
	}
	return d
}

// Close close dao.
func (dao *Dao) Close() {
	if dao.db != nil {
		dao.db.Close()
	}
}

// Ping ping cpdb
func (dao *Dao) Ping(c context.Context) (err error) {
	return dao.db.Ping(c)
}

func (dao *Dao) InitMember(ctx context.Context, arg *model.Member) (err error) {
	return dao.dbInitMember(ctx, arg)
}

func (dao *Dao) AddMember(ctx context.Context, arg *model.Member) (id int64, err error) {
	return dao.dbAddMember(ctx, arg)
}

func (dao *Dao) BatchAddMember(ctx context.Context, args []*model.Member) (affectRow int64, err error) {
	return dao.dbBatchAddMember(ctx, args)
}

func (dao *Dao) GetMemberByID(ctx context.Context, id int64) (m *model.Member, err error) {
	return dao.dbGetMemberByID(ctx, id)
}

func (dao *Dao) GetMemberByPhone(ctx context.Context, phone string) (m *model.Member, err error) {
	return dao.dbGetMemberByPhone(ctx, phone)
}

func (dao *Dao) GetMemberMaxAge(ctx context.Context) (age int64, err error) {
	return dao.dbGetMemberMaxAge(ctx)
}

func (dao *Dao) GetMemberSumAge(ctx context.Context) (age int64, err error) {
	return dao.dbGetMemberSumAge(ctx)
}

func (dao *Dao) CountMember(ctx context.Context) (count int64, err error) {
	return dao.dbCountMember(ctx)
}

func (dao *Dao) ListMember(ctx context.Context) (res []*model.Member, err error) {
	return dao.dbListMember(ctx)
}

func (dao *Dao) HasMemberByID(ctx context.Context, id int64) (has bool, err error) {
	return dao.dbHasMemberByID(ctx, id)
}

func (dao *Dao) QueryMemberByName(ctx context.Context, name string) (res []*model.Member, err error) {
	return dao.dbQueryMemberByName(ctx, name)
}

func (dao *Dao) QueryMemberByIDs(ctx context.Context, ids []int64) (res map[int64]*model.Member, err error) {
	return dao.dbQueryMemberByIDs(ctx, ids)
}

func (dao *Dao) UpdateMember(ctx context.Context, member *model.Member) (err error) {
	return dao.dbUpdateMember(ctx, member)
}

func (dao *Dao) SetMember(ctx context.Context, arg *model.Member) (err error) {
	return dao.dbSetMember(ctx, arg)
}

func (dao *Dao) SortMember(ctx context.Context, args model.ArgMemberSort) (err error) {
	return dao.dbSortMember(ctx, args)
}

func (dao *Dao) DeleteMember(ctx context.Context, id int64) (err error) {
	return dao.dbDeleteMember(ctx, id)
}
