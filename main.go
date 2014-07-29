package main

import (
	"github.com/astaxie/beego"
	_ "github.com/piotrkowalczuk/gowik-tracker/routers"
	"github.com/piotrkowalczuk/gowik-tracker/services"
)

func main() {
	mongo := services.NewMongo("localhost", "gowik")
	mongo.Connect()

	beego.Run()
}
