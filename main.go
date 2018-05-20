package main

import (
	"fmt"

	"github.com/mgerb/spam-filter-bot/bot"
	"github.com/mgerb/spam-filter-bot/config"
	log "github.com/sirupsen/logrus"
)

var version = "undefined"

func init() {
	fmt.Println("Starting spam filter bot " + version)

	log.SetLevel(log.DebugLevel)
	config.Init()
}

func main() {
	bot.Init()
}
