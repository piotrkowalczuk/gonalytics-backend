package controllers

import (
	"net/http"
	"strconv"

	"labix.org/v2/mgo/bson"
)

// GeneralController contains common properties accross multiple controllers
type GeneralController struct {
	BaseController
	ResponseData interface{}
}

// Finish is called once the method completes
func (gc *GeneralController) Finish() {
	outputFormat := gc.GetString("format")

	if outputFormat != "xml" && outputFormat != "json" {
		gc.Ctx.Abort(http.StatusBadRequest, "Output format parameter is missing.")
	}

	gc.Data[outputFormat] = &gc.ResponseData

	if outputFormat == "xml" {
		gc.ServeXml()
	} else if outputFormat == "json" {
		gc.ServeJson()
	}
}

// GetQueryLimit ...
func (gc *GeneralController) GetQueryLimit() int {
	limit, err := strconv.ParseInt(gc.Ctx.Input.Query("limit"), 10, 32)

	if err != nil {
		return 0
	}

	return int(limit)
}

// GetQuerySkip ...
func (gc *GeneralController) GetQuerySkip() int {
	skip, err := strconv.ParseInt(gc.Ctx.Input.Query("offset"), 10, 32)

	if err != nil {
		return 0
	}

	return int(skip)
}

// GetQuerySelect ...
func (gc *GeneralController) GetQuerySelect() bson.M {
	fields := gc.GetStrings("fields")
	selectFields := bson.M{}

	for key := range fields {
		selectFields[fields[key]] = 1
	}

	return selectFields
}
