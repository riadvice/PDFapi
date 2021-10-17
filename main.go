package main

import (
	log "github.com/sirupsen/logrus"
	"pdfannotations/server"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	// Enable if debug
	// log.SetReportCaller(true)
	server.HandleRequests()
}
