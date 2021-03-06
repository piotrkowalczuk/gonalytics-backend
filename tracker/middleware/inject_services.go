package middleware

import (
	"github.com/gocraft/web"
	"github.com/piotrkowalczuk/gonalytics-backend/services"
	CtrlV1 "github.com/piotrkowalczuk/gonalytics-backend/tracker/controllers/v1"
)

// InjectServicesMiddleware ...
func InjectServicesMiddleware(bc *CtrlV1.BaseContext, rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	bc.Logger = services.Logger
	bc.RepositoryManager = services.RepositoryManager
	bc.KafkaPublisher = services.KafkaPublisher

	next(rw, req)
}
