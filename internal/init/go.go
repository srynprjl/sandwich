package init

import (
	"fmt"
	"os"
	"os/exec"
)

func InitGo(project_data map[string]any) {
	file := project_data["path"].(string) + ".sandwich_initialized"
	_, fileErr := os.Stat(file)
	if !os.IsNotExist(fileErr) {
		fmt.Println("Error: project already initialized")
		return
	}
	os.MkdirAll(project_data["path"].(string), 0700)
	_, err := exec.LookPath("go")
	if err != nil {
		fmt.Println("Golang not found in PATH!!! Please download the toolchain")
		return
	}
	var moduleName string
	fmt.Print("Enter the go module name: \n>> ")
	fmt.Scan(&moduleName)

	if moduleName == "" {
		fmt.Println("Error: Please input a module name")
		return
	}

	os.Chdir(project_data["path"].(string))
	os.Create(".sandwich_initialized")
	goMod := exec.Command("go", "mod", "init", moduleName)
	goErr := goMod.Run()
	if goErr != nil {
		fmt.Println("Error:", err)
	}
	os.MkdirAll("./cmd/"+project_data["name"].(string), 0700)
	fil, filErr := os.Create("./cmd/" + project_data["name"].(string) + "/main.go")
	if filErr != nil {
		fmt.Println("Error:", filErr.Error())
		return
	}
	// temporary solution until i use git to write templates
	fil.WriteString(`
			package main
			func main(){
				print("This project was created using Sandwich")
			}
			`)
	fil.Close()
	// Initialize a git repository

}
