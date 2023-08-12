package handler

import (
	"eos-layout/internal/dto/request"
	"eos-layout/internal/service"
	"eos-layout/internal/state"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AreaHandler interface {
	Province(ctx *gin.Context)
	ProvinceList(ctx *gin.Context)
	City(ctx *gin.Context)
	CityList(ctx *gin.Context)
	District(ctx *gin.Context)
	DistrictList(ctx *gin.Context)
	Street(ctx *gin.Context)
	StreetList(ctx *gin.Context)
	Committee(ctx *gin.Context)
	CommitteeList(ctx *gin.Context)
}

func NewAreaHandler(h *Handler, areaService service.AreaService) AreaHandler {
	return &areaHandler{h, areaService}
}

type areaHandler struct {
	*Handler
	areaService service.AreaService
}

func (h *areaHandler) Province(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		h.Error(ctx, state.ErrorInvalidParams)
		return
	}
	in := &request.Area{
		Level: 0,
		ID:    id,
	}
	res, err := h.areaService.One(ctx, in)
	if err != nil {
		h.Error(ctx, err)
		return
	}
	h.Success(ctx, res)
}

func (h *areaHandler) ProvinceList(ctx *gin.Context) {
	key := ctx.Query("key")
	in := &request.Area{
		Level: 0,
		ID:    0,
		Key:   key,
	}
	res, err := h.areaService.Find(ctx, in)
	if err != nil {
		h.Error(ctx, err)
		return
	}
	h.Success(ctx, res)
}

func (h *areaHandler) City(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		h.Error(ctx, state.ErrorInvalidParams)
		return
	}
	in := &request.Area{
		Level: 1,
		ID:    id,
	}
	res, err := h.areaService.One(ctx, in)
	if err != nil {
		h.Error(ctx, err)
		return
	}
	h.Success(ctx, res)
}

func (h *areaHandler) CityList(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		h.Error(ctx, state.ErrorInvalidParams)
		return
	}
	key := ctx.Query("key")
	in := &request.Area{
		Level: 1,
		ID:    id,
		Key:   key,
	}
	res, err := h.areaService.Find(ctx, in)
	if err != nil {
		h.Error(ctx, err)
		return
	}
	h.Success(ctx, res)
}

func (h *areaHandler) District(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		h.Error(ctx, state.ErrorInvalidParams)
		return
	}
	in := &request.Area{
		Level: 2,
		ID:    id,
	}
	res, err := h.areaService.One(ctx, in)
	if err != nil {
		h.Error(ctx, err)
		return
	}
	h.Success(ctx, res)
}

func (h *areaHandler) DistrictList(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		h.Error(ctx, state.ErrorInvalidParams)
		return
	}
	key := ctx.Query("key")
	in := &request.Area{
		Level: 2,
		ID:    id,
		Key:   key,
	}
	res, err := h.areaService.Find(ctx, in)
	if err != nil {
		h.Error(ctx, err)
		return
	}
	h.Success(ctx, res)
}

func (h *areaHandler) Street(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		h.Error(ctx, state.ErrorInvalidParams)
		return
	}
	in := &request.Area{
		Level: 4,
		ID:    id,
	}
	res, err := h.areaService.One(ctx, in)
	if err != nil {
		h.Error(ctx, err)
		return
	}
	h.Success(ctx, res)
}

func (h *areaHandler) StreetList(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		h.Error(ctx, state.ErrorInvalidParams)
		return
	}
	key := ctx.Query("key")
	in := &request.Area{
		Level: 4,
		ID:    id,
		Key:   key,
	}
	res, err := h.areaService.Find(ctx, in)
	if err != nil {
		h.Error(ctx, err)
		return
	}
	h.Success(ctx, res)
}

func (h *areaHandler) Committee(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		h.Error(ctx, state.ErrorInvalidParams)
		return
	}
	in := &request.Area{
		Level: 5,
		ID:    id,
	}
	res, err := h.areaService.One(ctx, in)
	if err != nil {
		h.Error(ctx, err)
		return
	}
	h.Success(ctx, res)
}

func (h *areaHandler) CommitteeList(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		h.Error(ctx, state.ErrorInvalidParams)
		return
	}
	key := ctx.Query("key")
	in := &request.Area{
		Level: 5,
		ID:    id,
		Key:   key,
	}
	res, err := h.areaService.Find(ctx, in)
	if err != nil {
		h.Error(ctx, err)
		return
	}
	h.Success(ctx, res)
}
