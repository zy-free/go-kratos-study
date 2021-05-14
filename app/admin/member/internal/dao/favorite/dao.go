package favorite

import (
	"context"

	"go-kartos-study/app/admin/member/conf"
	"go-kartos-study/app/admin/member/internal/model"
	"go-kartos-study/pkg/database/orm"

	"github.com/jinzhu/gorm"
)

// Dao is redis dao.
type Dao struct {
	c *conf.Config
	// db
	db *gorm.DB
}

// New new a dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:  c,
		db: orm.NewMySQL(c.ORM),
	}
	d.initORM()
	return d
}

func (dao *Dao) initORM() {
	dao.db.LogMode(true)
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
