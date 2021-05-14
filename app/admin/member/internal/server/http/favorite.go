package http

import (
	"go-kartos-study/app/admin/member/internal/model"
	bm "go-kartos-study/pkg/net/http/blademaster"
)

func addFavorite(c *bm.Context) {
	c.JSON("addFavorite",nil)
}

func getFavoriteByID(c *bm.Context) {
	arg := new(model.GetFavoriteByIDReq)
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(favSvc.GetFavoriteByID(c, arg))
}

