package http

import bm "go-kartos-study/pkg/net/http/blademaster"

func addFavorite(c *bm.Context) {
	c.JSON("addFavorite",nil)
}
