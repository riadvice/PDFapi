package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	PORT        string
	OUTPUT      string
	INPUT       string
	SCRIPT_PATH string
	FONT_PATH   string
	EVENTS      string
)

func init() {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	log.WithFields(log.Fields{"ConfigurationFile": v.ConfigFileUsed()}).Info("Reading configuration file")

	PORT = v.GetString("Port")
	OUTPUT = v.GetString("OutputPath")
	INPUT = v.GetString("BBBPresPath")
	SCRIPT_PATH = v.GetString("ScriptDir")
	EVENTS = v.GetString("EventsPath")
	FONT_PATH = v.GetString("FontPath")
}
