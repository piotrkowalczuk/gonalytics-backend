package routers

import (
	"github.com/gocraft/web"
	CtrlV1 "github.com/piotrkowalczuk/gonalytics-backend/tracker/controllers/v1"
	"github.com/piotrkowalczuk/gonalytics-backend/tracker/middleware"
)

// GetRouterV1 ...
func GetRouterV1() *web.Router {
	return web.New(CtrlV1.BaseContext{}).
		Middleware(middleware.ExecutionDurationMiddleware).
		Middleware(middleware.InjectServicesMiddleware).
		Get("/v1/visit", (*CtrlV1.BaseContext).VisitsGETHandler)
}
