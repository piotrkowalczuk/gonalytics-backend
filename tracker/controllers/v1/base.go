package v1

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gocraft/web"
	geoip2 "github.com/oschwald/geoip2-golang"
	"github.com/piotrkowalczuk/gonalytics-backend/lib"
)

// BaseContext contains common properties accross multiple handlers
type BaseContext struct {
	RepositoryManager lib.RepositoryManager
	KafkaPublisher    *lib.KafkaPublisher
	Logger            *logrus.Logger
	GeoIP             *geoip2.Reader
	Response          interface{}
}

// HTTPError ...
func (bc *BaseContext) HTTPError(rw web.ResponseWriter, err error, message string, code int) {
	bc.Logger.Error(err.Error())
	http.Error(rw, message, code)
}
