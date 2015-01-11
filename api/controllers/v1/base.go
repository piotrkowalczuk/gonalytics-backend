package v1

import (
	"encoding/json"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gocraft/web"
	"github.com/piotrkowalczuk/gonalytics-backend/lib"
)

const (
	// AggregateByDay ...
	AggregateByDay = "day"
	// AggregateByMonth ...
	AggregateByMonth = "month"
	// AggregateByYear ...
	AggregateByYear = "year"
)

// BaseContext contains common properties accross multiple handlers
type BaseContext struct {
	RepositoryManager lib.RepositoryManager
	Logger            *logrus.Logger
	Response          interface{}
}

// HTTPError ...
func (bc *BaseContext) HTTPError(rw web.ResponseWriter, err error, message string, code int) {
	bc.Logger.Error(err.Error())
	http.Error(rw, message, code)
}

// ServeJSON ...
func (bc *BaseContext) ServeJSON(rw web.ResponseWriter, result interface{}) {
	js, err := json.Marshal(result)
	if err != nil {
		bc.HTTPError(rw, err, "Unexpected error.", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)
}
