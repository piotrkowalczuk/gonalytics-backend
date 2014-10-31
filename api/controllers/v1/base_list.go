package v1

import (
	"strconv"
)

// BaseListController contains common properties accross multiple controllerss
type BaseListController struct {
	BaseController
}

// GetQueryLimit ...
func (blc *BaseListController) GetQueryLimit() int {
	limit, err := strconv.ParseInt(blc.Ctx.Input.Query("limit"), 10, 32)

	if err != nil {
		return 0
	}

	return int(limit)
}

// GetQuerySkip ...
func (blc *BaseListController) GetQuerySkip() int {
	skip, err := strconv.ParseInt(blc.Ctx.Input.Query("offset"), 10, 32)

	if err != nil {
		return 0
	}

	return int(skip)
}

// GetQuerySelect ...
func (blc *BaseListController) GetQuerySelect() cql.M {
	fields := gc.GetStrings("fields")
	selectFields := cql.M{}

	for key := range fields {
		selectFields[fields[key]] = 1
	}

	return selectFields
}
