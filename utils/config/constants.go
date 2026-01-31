package config

import (
	"fmt"
	"log"
	"os"
)

var Conf Config
var PROJECT_NAME = "sandwich"
var lol, _ = os.UserConfigDir()
var CONFIG_LOCATION = fmt.Sprintf("%s/%s/", lol, PROJECT_NAME)

func InitializeConfig() {
	var err error
	Conf, err = LoadConfig(CONFIG_LOCATION)
	if err != nil {
		log.Fatal(err)
	}
}
