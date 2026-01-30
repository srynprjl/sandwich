package utils

import (
	"fmt"
	"log"
	"os"
)

var Conf Config
var USER = os.Getenv("USER")
var PROJECT_NAME = "sandwich"
var CONFIG_LOCATION = fmt.Sprintf("/home/%s/.config/%s/", USER, PROJECT_NAME)

func InitializeVariables() {
	c, err := LoadConfig(CONFIG_LOCATION)
	if err != nil {
		log.Fatal(err)
	}
	Conf = c
}
