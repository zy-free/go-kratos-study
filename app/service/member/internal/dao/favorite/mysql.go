package favorite

import (
	"context"
	"github.com/pkg/errors"
	"go-kartos-study/app/service/member/internal/model"
)

// 根据id查询单个
func (dao *Dao) dbGetFavoriteByID(ctx context.Context, id int64) (m *model.Favorite, err error) {
	m = &model.Favorite{}
	_sql := `SELECT id,mid,name,hint_at FROM member_favorite WHERE id = ? `
	if err = dao.db.QueryRow(ctx, _sql, id).Scan(&m.Id, &m.Mid, &m.Name, &m.HintAt); err != nil {
		return nil, errors.Wrapf(err, "dbGetFavoriteByID id(%d)", id)
	}
	return
}
