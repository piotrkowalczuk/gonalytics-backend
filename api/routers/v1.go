package routers

import (
	"github.com/gocraft/web"
	CtrlV1 "github.com/piotrkowalczuk/gonalytics-backend/api/controllers/v1"
	"github.com/piotrkowalczuk/gonalytics-backend/api/middleware"
)

func GetRouterV1() *web.Router {
	return web.New(CtrlV1.BaseContext{}).
		Middleware(middleware.ExecutionDurationMiddleware).
		Middleware(middleware.InjectServicesMiddleware).
		Get("/v1/sites/:siteId/nb-of-actions-by-country", (*CtrlV1.BaseContext).NbOfActionsByCountryGETHandler).
		Get("/v1/sites/:siteId/nb-of-actions-by-browser", (*CtrlV1.BaseContext).NbOfActionsByBrowserGETHandler).
		Get("/v1/sites/:siteId/nb-of-visits-by-country", (*CtrlV1.BaseContext).NbOfVisitsByCountryGETHandler).
		Get("/v1/sites/:siteId/nb-of-visits-by-browser", (*CtrlV1.BaseContext).NbOfVisitsByBrowserGETHandler)
}
