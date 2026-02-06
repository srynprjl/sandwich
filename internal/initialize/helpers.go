package initialize

import (
	"fmt"
	"os"
	"os/exec"
)

func checkProjectInitiated(path string) bool {
	file := path + ".stack"
	_, fileErr := os.Stat(file)
	if !os.IsNotExist(fileErr) {
		return true
	}
	os.MkdirAll(path, 0700)
	return false
}

func checkDependencies(dependencies ...string) bool {
	flag := true
	for _, dependency := range dependencies {
		_, err := exec.LookPath(dependency)
		if err != nil {
			fmt.Printf("%s not found in path. Please download it.", dependency)
			flag = false
		}
	}
	return flag
}

func runCommand(cmd string, args ...string) error {
	com := exec.Command(cmd, args...)
	err := com.Run()
	return err
}
