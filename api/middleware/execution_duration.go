package middleware

import (
	"bytes"
	"strconv"
	"time"

	"github.com/gocraft/web"
	CtrlV1 "github.com/piotrkowalczuk/gonalytics-backend/api/controllers/v1"
)

// ExecutionDurationMiddleware ...
func ExecutionDurationMiddleware(bc *CtrlV1.BaseContext, rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	startedAt := time.Now()
	next(rw, req)

	b := bytes.NewBufferString("[")
	b.WriteString(req.Method)
	b.WriteString("] ")
	b.WriteString("[")
	b.WriteString(strconv.FormatInt(int64(rw.StatusCode()), 10))
	b.WriteString("] ")
	b.WriteString(req.URL.Path)
	b.WriteString(" ")
	b.WriteString(time.Since(startedAt).String())

	bc.Logger.Info(b.String())
}
