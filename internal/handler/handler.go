package handler

import (
	"eos-layout/internal/state"
	"eos-layout/pkg/log"

	"github.com/gin-gonic/gin"
)

func NewHandler(log *log.Logger) *Handler {
	return &Handler{log: log}
}

type Handler struct {
	log *log.Logger
}

type response struct {
	state.Response
	Data any `json:"data"`
}

func (h *Handler) Success(ctx *gin.Context, data any) {
	if data == nil {
		data = map[string]string{}
	}
	resp := response{Data: data}
	resp.Code = 0
	resp.Message = "success"
	ctx.JSON(state.Success.StatusCode(), resp)
}

func (h *Handler) Error(ctx *gin.Context, err error) {
	switch e := err.(type) {
	case *state.Error:
		ctx.JSON(e.Response())
	default:
		ctx.JSON(state.ErrorServer.Response(e))
	}
}
