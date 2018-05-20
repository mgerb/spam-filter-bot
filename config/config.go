package config

import (
	"encoding/json"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

var (
	Config configFile
)

type configFile struct {
	Token string `json:"token"`
	// channel id to post filtered content
	FilterChannelID string `json:"filter_channel_id"`
	// role to apply filter to
	FilterRoleID string `json:"filter_role_id"`
}

// Init - read config file
func Init() {
	parseConfig()
}

func parseConfig() {
	log.Debug("Reading config file...")

	file, e := ioutil.ReadFile("./config.json")

	if e != nil {
		log.Fatal("File error: %v\n", e)
	}

	log.Debug("%s\n", string(file))

	err := json.Unmarshal(file, &Config)

	if err != nil {
		log.Error(err)
	}
}
