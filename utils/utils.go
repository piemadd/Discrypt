package utils

import (
	"log"
)

func HandleFatal(loc string, err error) {
	if err != nil {
		log.Fatal("Fatal error at " + loc + "\n\n" + err.Error())
	}
}
