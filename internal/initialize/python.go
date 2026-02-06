package initialize

import (
	"fmt"
	"os"
)

// TODO:
// way to initialize a python project (for fastAPI, data science, flask, django, etc )

func InitPython(projectData map[string]any, dependencies []string) {
	initiated := checkProjectInitiated(projectData["path"].(string))
	if initiated {
		fmt.Println("projected already initiated.")
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
	os.Create(".stack")
	err := runCommand("uv", "init")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = runCommand("uv", "venv")
	if err != nil {
		fmt.Println(err.Error())
	}
	if len(dependencies) != 0 {
		fmt.Println("Installing dependencies... ")
		args := []string{"add"}
		args = append(args, dependencies...)
		err = runCommand("uv", args...)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	fmt.Println("Project Initiated. ")
}
