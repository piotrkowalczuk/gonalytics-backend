package v1

import (
	"github.com/Sirupsen/logrus"
	"github.com/gocraft/web"
	geoip2 "github.com/oschwald/geoip2-golang"
	"github.com/piotrkowalczuk/gonalytics-backend/lib"
	"net/http"
)

// BaseController contains common properties accross multiple handlers
type BaseContext struct {
	RepositoryManager lib.RepositoryManager
	Logger            *logrus.Logger
	GeoIP             *geoip2.Reader
	Response          interface{}
}

func (bc *BaseContext) HTTPError(rw web.ResponseWriter, err error, message string, code int) {
	bc.Logger.Error(err.Error())
	http.Error(rw, message, code)
}
