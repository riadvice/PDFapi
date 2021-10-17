package main

import (
	log "github.com/sirupsen/logrus"
	"pdfannotations/server"

	log "github.com/sirupsen/logrus"
)

func main() {
<<<<<<< HEAD
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	// Enable if debug
	log.SetReportCaller(false)
||||||| parent of bd6f49b (- Add Vagrant configuration for dev.)
=======
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	// Enable if debug
	// log.SetReportCaller(true)
>>>>>>> bd6f49b (- Add Vagrant configuration for dev.)
	server.HandleRequests()
}
