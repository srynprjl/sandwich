package utils

import (
	"fmt"
	"os"
)

const PROJECT_NAME = "sandwich"
const AUTHOR = "sysnefo"
const DATABASE_NAME = "sandwich_go"

var USER = os.Getenv("USER")
var DATABASE_LOCATION = fmt.Sprintf("/home/%s/.local/share/%s/%s/", USER, AUTHOR, PROJECT_NAME)
