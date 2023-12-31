package handler

import (
	"eos-layout/internal/status"
	"eos-layout/pkg/log"
	"eos-layout/pkg/verifier"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	NewHandler,
	NewAreaHandler,
)

type Handler struct {
	log *log.Logger
}

func NewHandler(log *log.Logger) *Handler {
	return &Handler{log: log}
}

type response struct {
	status.Response
	Data any `json:"data"`
}

func (h *Handler) Log() *log.Logger {
	return h.log
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
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
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
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
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

func NoMethodHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(405, gin.H{"code": 405, "message": "methods are not allowed"})
	}
}

func NoRouteHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(500, gin.H{"code": 500, "message": "no handler function was found for the request route"})
	}
}
