package main

import (
	"flag"
	"os"
	"runtime"

	"github.com/fatih/color"

	"github.com/astaxie/beego"
	_ "github.com/piotrkowalczuk/gonalytics-tracker/routers"
	"github.com/piotrkowalczuk/gonalytics-tracker/services"
)

type flags struct {
	runTracker bool
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var flags = new(flags)
	flag.Usage = flagUsage

	flag.BoolVar(&flags.runTracker, "run:tracker", false, "Run tracker.")
	flag.Parse()

	if flags.runTracker {
		runTracker()
	} else {
		flag.Usage()
	}
}

func runTracker() {
	services.InitLogger()
	mongoDB := services.InitMongoDB("mongodb://mongodb/gonalytics")
	services.InitRepositoryManager(mongoDB)

	beego.Run()
}
func flagUsage() {
	color.Green("Usage of Gonalytics:\n")
	flag.PrintDefaults()
	os.Exit(2)
}
