package main

import (
	"pdfannotations/server"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	// Enable if debug
	log.SetReportCaller(false)
	server.HandleRequests()
}
