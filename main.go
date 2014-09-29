package main

import (
	"runtime"

	_ "github.com/piotrkowalczuk/gonalytics-backend/cli"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
