package middleware

import (
	"bytes"
	"encoding/json"
	"eos-layout/pkg/log"
	"eos-layout/pkg/md5"
	"eos-layout/pkg/uuid"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"strconv"
	"time"
)

func RequestLogger(log *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trace := md5.New(uuid.UUID())
		log.NewContext(ctx, zap.String("trace", trace))
		log.NewContext(ctx, zap.String("method", ctx.Request.Method))
		headers, _ := json.Marshal(ctx.Request.Header)
		log.NewContext(ctx, zap.String("headers", string(headers)))
		log.NewContext(ctx, zap.String("url", ctx.Request.URL.String()))
		if ctx.Request.Body != nil {
			body, _ := ctx.GetRawData()
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			log.NewContext(ctx, zap.String("body", string(body)))
		}
		log.WithContext(ctx).Info("REQUEST")
		ctx.Next()
	}
}

func ResponseLogger(log *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rw := &responseWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = rw
		startTime := time.Now()
		ctx.Next()
		duration := int(time.Since(startTime).Milliseconds())
		ctx.Header("X-Response-Time", strconv.Itoa(duration))
		log.WithContext(ctx).Info("RESPONSE", zap.Any("body", rw.body.String()), zap.Any("time", fmt.Sprintf("%sms", strconv.Itoa(duration))))
	}
}

type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
