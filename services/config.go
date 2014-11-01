package services

import (
	"github.com/piotrkowalczuk/gonalytics-backend/lib"

	"encoding/xml"
	"log"
	"os"
)

// Config ...
var Config *lib.Config

// InitConfig ...
func InitConfig(environment string) {
	file := openFile("conf/" + environment + ".xml")
	defer file.Close()

	decoder := xml.NewDecoder(file)
	decoder.Decode(&Config)

	return
}

func openFile(filePath string) (file *os.File) {
	file, err := os.Open(filePath)

	if err != nil {
		log.Printf("Cannot open configuration file: %v\n", err)
		os.Exit(1)
	}

	return
}
