package controllers

import (
    "github.com/astaxie/beego"
    "github.com/piotrkowalczuk/gonalytics/services"
)

type (
    BaseController struct {
        beego.Controller
        MongoPool services.Pool
    }
)

// Prepare is called prior to the baseController method
func (bc *BaseController) Prepare() {
    bc.MongoPool = services.MongoPool
}

// Finish is called once the baseController method completes
func (bc *BaseController) Finish() {

}
