package initialize

import (
	"fmt"
	"os"
	"os/exec"
)

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
	uv := exec.Command("uv", "init")
	err := uv.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
	uv = exec.Command("uv", "venv")
	err = uv.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
	if len(dependencies) != 0 {
		fmt.Println("Installing dependencies... ")
		args := []string{"add"}
		args = append(args, dependencies...)
		uv = exec.Command("uv", args...)
		err = uv.Run()
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	fmt.Println("Project Initiated. ")
}
