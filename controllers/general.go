package controllers

import "net/http"

// GeneralController contains common properties accross multiple controllers
type GeneralController struct {
	BaseController
	ResponseData interface{}
}

// Finish is called once the method completes
func (gc *GeneralController) Finish() {
	outputFormat := gc.GetString("outputFormat")

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
