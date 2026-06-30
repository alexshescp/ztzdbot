// cmd/bot/logging.go
package main

import (
	"log"
	"os"
)

func InitLogger(debug bool) {
	log.SetOutput(os.Stdout)

	if debug {
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		log.Println("Logging: DEBUG mode enabled")
	} else {
		log.SetFlags(log.Ldate | log.Ltime)
	}
}
