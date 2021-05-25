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

func (dao *Dao) UpdateMember(ctx context.Context, arg *model.UpdateMemberReq) (err error) {
	return dao.dbUpdateMember(ctx, arg)
}
