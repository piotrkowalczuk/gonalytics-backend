package routers

import (
	"github.com/gocraft/web"
	CtrlV1 "github.com/piotrkowalczuk/gonalytics-backend/api/controllers/v1"
	"github.com/piotrkowalczuk/gonalytics-backend/api/middleware"
)

func GetRouterV1() *web.Router {
	return web.New(CtrlV1.BaseContext{}).
		Middleware(middleware.ExecutionDurationMiddleware).
		Middleware(middleware.InjectServicesMiddleware)
}
