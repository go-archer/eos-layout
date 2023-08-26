package handler

import (
	"eos-layout/internal/status"
	"eos-layout/pkg/log"
	"eos-layout/pkg/verifier"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func NewHandler(log *log.Logger) *Handler {
	return &Handler{log: log}
}

type Handler struct {
	log *log.Logger
}

type response struct {
	status.Response
	Data any `json:"data"`
}

func (h *Handler) Success(ctx *gin.Context, data any) {
	if data == nil {
		data = map[string]string{}
	}
	resp := response{Data: data}
	resp.Code = 0
	resp.Message = "success"
	ctx.JSON(status.Success.StatusCode(), resp)
}

func (h *Handler) Error(ctx *gin.Context, err error) {
	var e *status.Error
	switch {
	case errors.As(err, &e):
		ctx.JSON(e.Response())
	default:
		ctx.JSON(status.ErrorServer.Response(err))
	}
}

func (h *Handler) Bind(ctx *gin.Context, v any) error {
	err := ctx.ShouldBind(v)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return status.ErrorInvalidParams
		}
		return errs
	}
	return nil
}

func (h *Handler) Struct(v any) error {
	err := verifier.Validate.Struct(v)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return err
		}
		return errs
	}
	return nil
}

func (h *Handler) Var(v any, tag string, label ...string) error {
	err := verifier.Validate.Var(v, tag)
	if err != nil {
		res := verifier.Translate(err)
		msg := res
		if !verifier.IsEN() {
			if len(label) > 0 {
				msg = label[0] + msg
			}
		} else {
			if len(label) > 1 {
				msg = label[1] + msg
			}
		}
		return status.ErrorInvalidParams.Message(msg)
	}
	return nil
}
