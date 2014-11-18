package middleware

import (
	"bytes"
	"github.com/gocraft/web"
	CtrlV1 "github.com/piotrkowalczuk/gonalytics-backend/tracker/controllers/v1"
	"time"
)

// Which is equivalent to this:
func ExecutionDurationMiddleware(bc *CtrlV1.BaseContext, rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	startedAt := time.Now()
	next(rw, req)

	b := bytes.NewBufferString("[")
	b.WriteString(req.Method)
	b.WriteString("] ")
	b.WriteString(req.URL.Path)
	b.WriteString(" ")
	b.WriteString(time.Since(startedAt).String())

	bc.Logger.Info(b.String())
}
