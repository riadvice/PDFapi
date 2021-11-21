package main

import (
	log "github.com/sirupsen/logrus"
	"pdfannotations/server"

	log "github.com/sirupsen/logrus"
)

func main() {
<<<<<<< HEAD
<<<<<<< HEAD
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	// Enable if debug
	log.SetReportCaller(false)
||||||| parent of bd6f49b (- Add Vagrant configuration for dev.)
=======
||||||| parent of ddf7b04 (config + post + write text + arabic + optimization)
=======
<<<<<<< HEAD
>>>>>>> ddf7b04 (config + post + write text + arabic + optimization)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	// Enable if debug
	// log.SetReportCaller(true)
<<<<<<< HEAD
>>>>>>> bd6f49b (- Add Vagrant configuration for dev.)
||||||| parent of ddf7b04 (config + post + write text + arabic + optimization)
=======
||||||| parent of 2c42ff1 (config + post + write text + arabic + optimization)
=======
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	// Enable if debug
	log.SetReportCaller(false)
>>>>>>> 2c42ff1 (config + post + write text + arabic + optimization)
>>>>>>> ddf7b04 (config + post + write text + arabic + optimization)
	server.HandleRequests()
}
