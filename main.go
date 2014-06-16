package main

import (
	"github.com/astaxie/beego"
	_ "github.com/piotrkowalczuk/gonalytics/routers"
	"github.com/piotrkowalczuk/gonalytics/services"
)

func main() {
	mongo := services.NewMongo("localhost", "gowik")
	mongo.Connect()
	beego.Run()
}
