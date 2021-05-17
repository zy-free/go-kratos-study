package favorite

import (
	"context"

	"github.com/jinzhu/gorm"
	"go-kartos-study/app/admin/member/internal/model"
)

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

func (dao *Dao) GetFavoriteByID(ctx context.Context, id int64) (m *model.Favorite, err error) {
	return dao.dbGetFavoriteByID(ctx, id)
}
