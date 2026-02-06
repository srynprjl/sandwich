package initialize

import (
	"fmt"
	"os"
	"os/exec"
)

/*
 * src/main.py
 * pyproject.toml
 */

func InitPython(projectData map[string]any) {
	initiated := checkProjectInitiated(projectData["path"].(string))
	if initiated {
		fmt.Println("projected initiated.")
		return
	}
	pythonExists := checkDependencies("python")
	if !pythonExists {
		return
	}

	uvExists := checkDependencies("uv")
	if !uvExists {
		return
	}
	os.Chdir(projectData["path"].(string))
	uv := exec.Command("uv", "init")
	err := uv.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Project Initiated. ")

	// ask if you want to initialize a virtual environment
	// 
}
