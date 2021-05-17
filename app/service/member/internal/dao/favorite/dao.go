package favorite

import (
	"context"
	"go-kartos-study/app/service/member/internal/model"
	"go-kartos-study/pkg/database/sql"
)

// Dao is redis dao.
type Dao struct {
	// db
	db *sql.DB
}

// New new a dao.
func New(db *sql.DB) (d *Dao) {
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
	return dao.db.Ping(c)
}

func (dao *Dao) GetFavoriteByID(ctx context.Context, id int64) (m *model.Favorite, err error) {
	return dao.dbGetFavoriteByID(ctx, id)
}
