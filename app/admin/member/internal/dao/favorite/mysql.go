package favorite

import (
	"context"
	"github.com/pkg/errors"
	"go-kartos-study/app/admin/member/internal/model"
	"go-kartos-study/pkg/ecode"
)

// 根据id查询单个
func (dao *Dao) dbGetFavoriteByID(ctx context.Context, id int64) (m *model.Favorite, err error) {
	m = new(model.Favorite)
	if err = dao.db.Select("id,mid,name,hint_at").Table("member_favorite").Where("id=?",id).First(m).Error;err != nil{
		if err == ecode.NothingFound{
			err = nil
			return
		}
		err = errors.Wrapf(err, "dbGetFavoriteByID")
	}
	return
}
