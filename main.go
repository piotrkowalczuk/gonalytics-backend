package main

import (
	"github.com/astaxie/beego"
	_ "github.com/piotrkowalczuk/gonalytics-tracker/routers"
	"github.com/piotrkowalczuk/gonalytics-tracker/services"
)

func main() {
	mongo := services.NewMongo("localhost", "gowik")
	mongo.Connect()

	beego.Run()
}
