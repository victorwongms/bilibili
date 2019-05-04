// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: api.proto

/*
Package v1 is a generated blademaster stub package.
This code was generated with go-common/app/tool/bmgen/protoc-gen-bm v0.1.

It is generated from these files:
	api.proto
*/
package v1

import (
	"context"

	bm "go-common/library/net/http/blademaster"
	"go-common/library/net/http/blademaster/binding"
)

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathDaoAnchorFetchRoomByIDs = "/live.daoanchor.v1.DaoAnchor/FetchRoomByIDs"
var PathDaoAnchorRoomOnlineList = "/live.daoanchor.v1.DaoAnchor/RoomOnlineList"
var PathDaoAnchorRoomOnlineListByArea = "/live.daoanchor.v1.DaoAnchor/RoomOnlineListByArea"
var PathDaoAnchorRoomOnlineListByAttrs = "/live.daoanchor.v1.DaoAnchor/RoomOnlineListByAttrs"
var PathDaoAnchorRoomCreate = "/live.daoanchor.v1.DaoAnchor/RoomCreate"
var PathDaoAnchorRoomUpdate = "/live.daoanchor.v1.DaoAnchor/RoomUpdate"
var PathDaoAnchorRoomBatchUpdate = "/live.daoanchor.v1.DaoAnchor/RoomBatchUpdate"
var PathDaoAnchorRoomExtendUpdate = "/live.daoanchor.v1.DaoAnchor/RoomExtendUpdate"
var PathDaoAnchorRoomExtendBatchUpdate = "/live.daoanchor.v1.DaoAnchor/RoomExtendBatchUpdate"
var PathDaoAnchorRoomExtendIncre = "/live.daoanchor.v1.DaoAnchor/RoomExtendIncre"
var PathDaoAnchorRoomExtendBatchIncre = "/live.daoanchor.v1.DaoAnchor/RoomExtendBatchIncre"
var PathDaoAnchorRoomTagCreate = "/live.daoanchor.v1.DaoAnchor/RoomTagCreate"
var PathDaoAnchorRoomAttrCreate = "/live.daoanchor.v1.DaoAnchor/RoomAttrCreate"
var PathDaoAnchorRoomAttrSetEx = "/live.daoanchor.v1.DaoAnchor/RoomAttrSetEx"
var PathDaoAnchorAnchorUpdate = "/live.daoanchor.v1.DaoAnchor/AnchorUpdate"
var PathDaoAnchorAnchorBatchUpdate = "/live.daoanchor.v1.DaoAnchor/AnchorBatchUpdate"
var PathDaoAnchorAnchorIncre = "/live.daoanchor.v1.DaoAnchor/AnchorIncre"
var PathDaoAnchorAnchorBatchIncre = "/live.daoanchor.v1.DaoAnchor/AnchorBatchIncre"
var PathDaoAnchorFetchAreas = "/live.daoanchor.v1.DaoAnchor/FetchAreas"
var PathDaoAnchorFetchAttrByIDs = "/live.daoanchor.v1.DaoAnchor/FetchAttrByIDs"
var PathDaoAnchorDeleteAttr = "/live.daoanchor.v1.DaoAnchor/DeleteAttr"

// ===================
// DaoAnchor Interface
// ===================

type DaoAnchorBMServer interface {
	// FetchRoomByIDs 查询房间信息
	FetchRoomByIDs(ctx context.Context, req *RoomByIDsReq) (resp *RoomByIDsResp, err error)

	// RoomOnlineList 在线房间列表
	RoomOnlineList(ctx context.Context, req *RoomOnlineListReq) (resp *RoomOnlineListResp, err error)

	// RoomOnlineListByArea 分区在线房间列表(只返回room_id列表，不传分区，默认查找所有)
	RoomOnlineListByArea(ctx context.Context, req *RoomOnlineListByAreaReq) (resp *RoomOnlineListByAreaResp, err error)

	// RoomOnlineListByAttrs 在线房间维度信息(不传attrs，不查询attr)
	RoomOnlineListByAttrs(ctx context.Context, req *RoomOnlineListByAttrsReq) (resp *RoomOnlineListByAttrsResp, err error)

	// RoomCreate 房间创建
	RoomCreate(ctx context.Context, req *RoomCreateReq) (resp *RoomCreateResp, err error)

	// RoomUpdate 房间信息更新
	RoomUpdate(ctx context.Context, req *RoomUpdateReq) (resp *UpdateResp, err error)

	// RoomBatchUpdate 房间信息批量更新
	RoomBatchUpdate(ctx context.Context, req *RoomBatchUpdateReq) (resp *UpdateResp, err error)

	// RoomExtendUpdate 房间扩展信息更新
	RoomExtendUpdate(ctx context.Context, req *RoomExtendUpdateReq) (resp *UpdateResp, err error)

	// RoomExtendBatchUpdate 房间扩展信息批量更新
	RoomExtendBatchUpdate(ctx context.Context, req *RoomExtendBatchUpdateReq) (resp *UpdateResp, err error)

	// RoomExtendIncre 房间信息增量更新
	RoomExtendIncre(ctx context.Context, req *RoomExtendIncreReq) (resp *UpdateResp, err error)

	// RoomExtendBatchIncre 房间信息批量增量更新
	RoomExtendBatchIncre(ctx context.Context, req *RoomExtendBatchIncreReq) (resp *UpdateResp, err error)

	// RoomTagCreate 房间Tag创建
	RoomTagCreate(ctx context.Context, req *RoomTagCreateReq) (resp *UpdateResp, err error)

	// RoomAttrCreate 房间Attr创建
	RoomAttrCreate(ctx context.Context, req *RoomAttrCreateReq) (resp *UpdateResp, err error)

	// RoomAttrSetEx 房间Attr更新
	RoomAttrSetEx(ctx context.Context, req *RoomAttrSetExReq) (resp *UpdateResp, err error)

	// AnchorUpdate 主播信息更新
	AnchorUpdate(ctx context.Context, req *AnchorUpdateReq) (resp *UpdateResp, err error)

	// AnchorBatchUpdate 主播信息批量更新
	AnchorBatchUpdate(ctx context.Context, req *AnchorBatchUpdateReq) (resp *UpdateResp, err error)

	// AnchorIncre 主播信息增量更新
	AnchorIncre(ctx context.Context, req *AnchorIncreReq) (resp *UpdateResp, err error)

	// AnchorBatchIncre 主播信息批量增量更新
	AnchorBatchIncre(ctx context.Context, req *AnchorBatchIncreReq) (resp *UpdateResp, err error)

	// FetchAreas 根据父分区号查询子分区
	FetchAreas(ctx context.Context, req *FetchAreasReq) (resp *FetchAreasResp, err error)

	// FetchAttrByIDs 批量根据房间号查询指标
	FetchAttrByIDs(ctx context.Context, req *FetchAttrByIDsReq) (resp *FetchAttrByIDsResp, err error)

	// DeleteAttr 删除某一个指标
	DeleteAttr(ctx context.Context, req *DeleteAttrReq) (resp *UpdateResp, err error)
}

var v1DaoAnchorSvc DaoAnchorBMServer

func daoAnchorFetchRoomByIDs(c *bm.Context) {
	p := new(RoomByIDsReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1DaoAnchorSvc.FetchRoomByIDs(c, p)
	c.JSON(resp, err)
}

func daoAnchorRoomOnlineList(c *bm.Context) {
	p := new(RoomOnlineListReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1DaoAnchorSvc.RoomOnlineList(c, p)
	c.JSON(resp, err)
}

func daoAnchorRoomOnlineListByArea(c *bm.Context) {
	p := new(RoomOnlineListByAreaReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1DaoAnchorSvc.RoomOnlineListByArea(c, p)
	c.JSON(resp, err)
}

func daoAnchorRoomOnlineListByAttrs(c *bm.Context) {
	p := new(RoomOnlineListByAttrsReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1DaoAnchorSvc.RoomOnlineListByAttrs(c, p)
	c.JSON(resp, err)
}

func daoAnchorRoomCreate(c *bm.Context) {
	p := new(RoomCreateReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1DaoAnchorSvc.RoomCreate(c, p)
	c.JSON(resp, err)
}

func daoAnchorRoomUpdate(c *bm.Context) {
	p := new(RoomUpdateReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1DaoAnchorSvc.RoomUpdate(c, p)
	c.JSON(resp, err)
}

func daoAnchorRoomBatchUpdate(c *bm.Context) {
	p := new(RoomBatchUpdateReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1DaoAnchorSvc.RoomBatchUpdate(c, p)
	c.JSON(resp, err)
}

func daoAnchorRoomExtendUpdate(c *bm.Context) {
	p := new(RoomExtendUpdateReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1DaoAnchorSvc.RoomExtendUpdate(c, p)
	c.JSON(resp, err)
}

func daoAnchorRoomExtendBatchUpdate(c *bm.Context) {
	p := new(RoomExtendBatchUpdateReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1DaoAnchorSvc.RoomExtendBatchUpdate(c, p)
	c.JSON(resp, err)
}

func daoAnchorRoomExtendIncre(c *bm.Context) {
	p := new(RoomExtendIncreReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1DaoAnchorSvc.RoomExtendIncre(c, p)
	c.JSON(resp, err)
}

func daoAnchorRoomExtendBatchIncre(c *bm.Context) {
	p := new(RoomExtendBatchIncreReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1DaoAnchorSvc.RoomExtendBatchIncre(c, p)
	c.JSON(resp, err)
}

func daoAnchorRoomTagCreate(c *bm.Context) {
	p := new(RoomTagCreateReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1DaoAnchorSvc.RoomTagCreate(c, p)
	c.JSON(resp, err)
}

func daoAnchorRoomAttrCreate(c *bm.Context) {
	p := new(RoomAttrCreateReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1DaoAnchorSvc.RoomAttrCreate(c, p)
	c.JSON(resp, err)
}

func daoAnchorRoomAttrSetEx(c *bm.Context) {
	p := new(RoomAttrSetExReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1DaoAnchorSvc.RoomAttrSetEx(c, p)
	c.JSON(resp, err)
}

func daoAnchorAnchorUpdate(c *bm.Context) {
	p := new(AnchorUpdateReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1DaoAnchorSvc.AnchorUpdate(c, p)
	c.JSON(resp, err)
}

func daoAnchorAnchorBatchUpdate(c *bm.Context) {
	p := new(AnchorBatchUpdateReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1DaoAnchorSvc.AnchorBatchUpdate(c, p)
	c.JSON(resp, err)
}

func daoAnchorAnchorIncre(c *bm.Context) {
	p := new(AnchorIncreReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1DaoAnchorSvc.AnchorIncre(c, p)
	c.JSON(resp, err)
}

func daoAnchorAnchorBatchIncre(c *bm.Context) {
	p := new(AnchorBatchIncreReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1DaoAnchorSvc.AnchorBatchIncre(c, p)
	c.JSON(resp, err)
}

func daoAnchorFetchAreas(c *bm.Context) {
	p := new(FetchAreasReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1DaoAnchorSvc.FetchAreas(c, p)
	c.JSON(resp, err)
}

func daoAnchorFetchAttrByIDs(c *bm.Context) {
	p := new(FetchAttrByIDsReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1DaoAnchorSvc.FetchAttrByIDs(c, p)
	c.JSON(resp, err)
}

func daoAnchorDeleteAttr(c *bm.Context) {
	p := new(DeleteAttrReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1DaoAnchorSvc.DeleteAttr(c, p)
	c.JSON(resp, err)
}

// RegisterDaoAnchorBMServer Register the blademaster route
func RegisterDaoAnchorBMServer(e *bm.Engine, server DaoAnchorBMServer) {
	v1DaoAnchorSvc = server
	e.GET("/live.daoanchor.v1.DaoAnchor/FetchRoomByIDs", daoAnchorFetchRoomByIDs)
	e.GET("/live.daoanchor.v1.DaoAnchor/RoomOnlineList", daoAnchorRoomOnlineList)
	e.GET("/live.daoanchor.v1.DaoAnchor/RoomOnlineListByArea", daoAnchorRoomOnlineListByArea)
	e.GET("/live.daoanchor.v1.DaoAnchor/RoomOnlineListByAttrs", daoAnchorRoomOnlineListByAttrs)
	e.GET("/live.daoanchor.v1.DaoAnchor/RoomCreate", daoAnchorRoomCreate)
	e.GET("/live.daoanchor.v1.DaoAnchor/RoomUpdate", daoAnchorRoomUpdate)
	e.GET("/live.daoanchor.v1.DaoAnchor/RoomBatchUpdate", daoAnchorRoomBatchUpdate)
	e.GET("/live.daoanchor.v1.DaoAnchor/RoomExtendUpdate", daoAnchorRoomExtendUpdate)
	e.GET("/live.daoanchor.v1.DaoAnchor/RoomExtendBatchUpdate", daoAnchorRoomExtendBatchUpdate)
	e.GET("/live.daoanchor.v1.DaoAnchor/RoomExtendIncre", daoAnchorRoomExtendIncre)
	e.GET("/live.daoanchor.v1.DaoAnchor/RoomExtendBatchIncre", daoAnchorRoomExtendBatchIncre)
	e.GET("/live.daoanchor.v1.DaoAnchor/RoomTagCreate", daoAnchorRoomTagCreate)
	e.GET("/live.daoanchor.v1.DaoAnchor/RoomAttrCreate", daoAnchorRoomAttrCreate)
	e.GET("/live.daoanchor.v1.DaoAnchor/RoomAttrSetEx", daoAnchorRoomAttrSetEx)
	e.GET("/live.daoanchor.v1.DaoAnchor/AnchorUpdate", daoAnchorAnchorUpdate)
	e.GET("/live.daoanchor.v1.DaoAnchor/AnchorBatchUpdate", daoAnchorAnchorBatchUpdate)
	e.GET("/live.daoanchor.v1.DaoAnchor/AnchorIncre", daoAnchorAnchorIncre)
	e.GET("/live.daoanchor.v1.DaoAnchor/AnchorBatchIncre", daoAnchorAnchorBatchIncre)
	e.GET("/live.daoanchor.v1.DaoAnchor/FetchAreas", daoAnchorFetchAreas)
	e.GET("/live.daoanchor.v1.DaoAnchor/FetchAttrByIDs", daoAnchorFetchAttrByIDs)
	e.GET("/live.daoanchor.v1.DaoAnchor/DeleteAttr", daoAnchorDeleteAttr)
}
