package services

import (
	"github.com/piotrkowalczuk/gonalytics-backend/lib"

	"encoding/xml"
	"log"
	"os"
)

// APIConfig ...
var APIConfig *lib.APIConfig

// TrackerConfig ...
var TrackerConfig *lib.TrackerConfig

// ActionsWorkerConfig ...
var ActionsWorkerConfig *lib.ActionsWorkerConfig

// InitConfig ...
func InitConfig(consumer string, environment string) {
	file := openFile("conf/" + consumer + "/" + environment + ".xml")
	defer file.Close()
	decoder := xml.NewDecoder(file)

	switch consumer {
	case lib.APIConfigConsumer:
		decoder.Decode(&APIConfig)
	case lib.TrackerConfigConsumer:
		decoder.Decode(&TrackerConfig)
	case lib.ActionsWorkerConfigConsumer:
		decoder.Decode(&ActionsWorkerConfig)
	}

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
