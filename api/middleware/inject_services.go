package middleware

import (
	"github.com/gocraft/web"
	CtrlV1 "github.com/piotrkowalczuk/gonalytics-backend/api/controllers/v1"
	"github.com/piotrkowalczuk/gonalytics-backend/services"
)

// InjectServicesMiddleware ...
func InjectServicesMiddleware(bc *CtrlV1.BaseContext, rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	bc.Logger = services.Logger
	bc.RepositoryManager = services.RepositoryManager

	next(rw, req)
}
