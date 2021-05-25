package http

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"time"

	"go-kartos-study/app/bff/member/internal/model"
	"go-kartos-study/pkg/log"
	bm "go-kartos-study/pkg/net/http/blademaster"
)

func getMemberByID(c *bm.Context) {
	arg := new(model.GetMemberByIDReq)
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(memSvc.GetMemberByID(c, arg))
}

func getMemberByPhone(c *bm.Context) {
	arg := new(model.GetMemberByPhoneReq)
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(memSvc.GetMemberByPhone(c, arg))
}

func getMemberMaxAge(c *bm.Context) {
	c.JSON(memSvc.GetMemberMaxAge(c))
}

func queryMemberByName(c *bm.Context) {
	arg := new(model.QueryMemberByNameReq)
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(memSvc.QueryMemberByName(c, arg))
}

func queryMemberByIDs(c *bm.Context) {
	arg := new(model.QueryMemberByIdsReq)
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(memSvc.QueryMemberByIds(c, arg))
}

func listMember(c *bm.Context) {
	arg := new(model.ListMemberReq)
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(memSvc.ListMember(c, arg))
}

func exportMember(c *bm.Context) {
	resp, err := memSvc.ListMember(c, &model.ListMemberReq{Pn: 1, Ps: 10000})
	if err != nil {
		c.JSON(nil, err)
		return
	}

	data := make([][]string, 0, len(resp.Members)+1)
	data = append(data, model.MemberExportFields())
	for _, v := range resp.Members {
		data = append(data, v.ExportStrings())
	}
	buf := new(bytes.Buffer)

	// add utf bom
	if len(data) > 0 {
		buf.WriteString("\xEF\xBB\xBF")
	}

	csvWriter := csv.NewWriter(buf)
	err = csvWriter.WriteAll(data)
	if err != nil {
		c.JSON(nil, err)
		return
	}
	c.CSV(bm.CSVMsg{
		Content: buf.Bytes(),
		Title:   fmt.Sprintf("%s-%s", time.Now().Format("2006-01-02"), "members"),
	}, nil)
}

func addMember(c *bm.Context) {
	arg := new(model.AddMemberReq)
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(memSvc.AddMember(c, arg))
}

func initMember(c *bm.Context) {
	arg := new(model.InitMemberReq)
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(nil, memSvc.InitMember(c, arg))
}

func batchAddMember(c *bm.Context) {
	arg := new(model.BatchAddMemberReq)
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(nil, memSvc.BatchAddMember(c, arg))
}

func updateMember(c *bm.Context) {
	arg := new(model.UpdateMemberReq)
	if err := c.Bind(arg); err != nil {
		return
	}
	log.Info("updateMember(%v)", arg)
	c.JSON(nil, memSvc.UpdateMember(c, arg))
}

func setMember(c *bm.Context) {
	arg := new(model.SetMemberReq)
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(nil, memSvc.SetMember(c, arg))
}

func sortMember(c *bm.Context) {
	arg := new(model.SortMemberReq)
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(nil, memSvc.SortMember(c, arg))
}

func delMember(c *bm.Context) {
	arg := new(model.DelMemberReq)
	if err := c.Bind(arg); err != nil {
		return
	}
	c.JSON(nil, memSvc.DelMember(c, arg))
}
