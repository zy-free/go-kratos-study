package member

import (
	"context"
	"github.com/jinzhu/gorm"
	"go-kartos-study/app/admin/member/internal/model"
)

// Dao is redis dao.
type Dao struct {
	db *gorm.DB
}

// New new a dao.
func New(db *gorm.DB) (d *Dao) {
	d = &Dao{
		db: db,
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
	return dao.db.DB().PingContext(c)
}

func (dao *Dao) AddMember(ctx context.Context, arg *model.AddMemberReq) (id int64, err error) {
	return dao.dbAddMember(ctx, arg)
}

func (dao *Dao) BatchAddMember(ctx context.Context, args []*model.AddMemberReq) (affectRow int64, err error) {
	return dao.dbBatchAddMember(ctx, args)
}

func (dao *Dao) GetMemberByID(ctx context.Context, id int64) (res *model.Member, err error) {
	return dao.dbGetMemberByID(ctx, id)
}

//func (dao *Dao) GetMemberByPhone(ctx context.Context, phone string) (m *model.Member, err error) {
//	return dao.dbGetMemberByPhone(ctx, phone)
//}
//
//func (dao *Dao) GetMemberMaxAge(ctx context.Context) (age int64, err error) {
//	return dao.dbGetMemberMaxAge(ctx)
//}
//
//func (dao *Dao) GetMemberSumAge(ctx context.Context) (age int64, err error) {
//	return dao.dbGetMemberSumAge(ctx)
//}
//
//func (dao *Dao) CountMember(ctx context.Context) (count int64, err error) {
//	return dao.dbCountMember(ctx)
//}
//
//func (dao *Dao) ListMember(ctx context.Context) (res []*model.Member, err error) {
//	return dao.dbListMember(ctx)
//}
//
//func (dao *Dao) HasMemberByID(ctx context.Context, id int64) (has bool, err error) {
//	return dao.dbHasMemberByID(ctx, id)
//}
//
//func (dao *Dao) QueryMemberByName(ctx context.Context, name string) (res []*model.Member, err error) {
//	return dao.dbQueryMemberByName(ctx, name)
//}
//
//func (dao *Dao) QueryMemberByIDs(ctx context.Context, ids []int64) (res map[int64]*model.Member, err error) {
//	return dao.dbQueryMemberByIDs(ctx, ids)
//}

func (dao *Dao) UpdateMember(ctx context.Context, arg *model.UpdateMemberReq) (err error) {
	return dao.dbUpdateMember(ctx, arg)
}

//func (dao *Dao) UpdateMemberAttr(ctx context.Context, id int64, attr int32) (err error) {
//	return dao.dbUpdateMemberAttr(ctx, id, attr)
//}
//
//func (dao *Dao) SetMember(ctx context.Context, arg *model.Member) (err error) {
//	return dao.dbSetMember(ctx, arg)
//}
//
//func (dao *Dao) SortMember(ctx context.Context, args model.ArgMemberSort) (err error) {
//	return dao.dbSortMember(ctx, args)
//}
//
//func (dao *Dao) DeleteMember(ctx context.Context, id int64) (err error) {
//	return dao.dbDeleteMember(ctx, id)
//}
