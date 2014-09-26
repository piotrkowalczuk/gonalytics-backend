package main

import (
	"runtime"

	"github.com/astaxie/beego"
	_ "github.com/piotrkowalczuk/gonalytics-backend/api/routers"
	"github.com/piotrkowalczuk/gonalytics-backend/service"
	_ "github.com/piotrkowalczuk/gonalytics-backend/tracker/routers"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	runTracker()
}

func runTracker() {
	service.InitLogger()
	mongoDB := service.InitMongoDB("mongodb://mongodb/gonalytics")
	cassandra := service.InitCassandra("gonalytics", []string{"127.0.0.1"})
	service.InitRepositoryManager(mongoDB, cassandra)

	defer cassandra.Close()
	beego.Run()
}
