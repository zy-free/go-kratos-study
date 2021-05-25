package http

import (
	"fmt"
	"go-kartos-study/app/admin/member/internal/model"
	"go-kartos-study/pkg/log"
	bm "go-kartos-study/pkg/net/http/blademaster"
)

func getMemberByID(c *bm.Context) {
	arg := new(model.GetMemberByIDReq)
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(memSvc.GetMemberByID(c, arg.ID))
}

func addMember(c *bm.Context) {
	arg := new(model.AddMemberReq)
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(memSvc.AddMember(c, arg))
}

func batchAddMember(c *bm.Context) {
	arg := new(model.BatchAddMemberReq)
	if err := c.Bind(arg); err != nil {
		return
	}
	fmt.Println(arg)
	_, err := memSvc.BatchAddMember(c, arg.Args)
	c.JSON(nil, err)
}

func updateMember(c *bm.Context) {
	arg := new(model.UpdateMemberReq)
	if err := c.Bind(arg); err != nil {
		return
	}
	log.Info("updateMember(%v)", arg)
	c.JSON(nil, memSvc.UpdateMember(c, arg))
}
