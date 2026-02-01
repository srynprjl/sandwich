package config

import (
	"fmt"
	"log"
	"os"
)

const VERSION = "1.0.0a"
const PROJECT_NAME = "sandwich"

var Conf Config
var CONFIG_DIR, _ = os.UserConfigDir()
var CONFIG_LOCATION = fmt.Sprintf("%s/%s/", CONFIG_DIR, PROJECT_NAME)

func InitializeConfig() {
	var err error
	Conf, err = LoadConfig(CONFIG_LOCATION)
	if err != nil {
		log.Fatal(err)
	}
}
