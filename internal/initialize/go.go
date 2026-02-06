package initialize

import (
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"path"
)

const goMainTemplate = `package main
import "fmt"
func main() {
	fmt.Println("This project was created using Sandwich")
}
`

// TODO:
// basic templates such as Gin, Gorilla Mux etc

func InitGo(project_data map[string]any) {
	projectPath := path.Join("cmd", project_data["name"].(string))
	filePath := path.Join(projectPath, "main.go")

	initiated := checkProjectInitiated(project_data["path"].(string))
	if initiated {
		fmt.Println("The project files already exist!!! ")
		return
	}
	dependencies := checkDependencies("go")
	if !dependencies {
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
	goMod := exec.Command("go", "mod", "init", moduleName)
	goErr := goMod.Run()
	if goErr != nil {
		fmt.Println("Error:", goErr)
	}
	os.Create(".stack")
	os.MkdirAll(projectPath, 0700)
	fil, filErr := os.Create(filePath)
	if filErr != nil {
		fmt.Println("Error:", filErr.Error())
		return
	}
	tmpl, err := template.New("main").Parse(goMainTemplate)
	if err != nil {
		fmt.Printf("Template error: %v\n", err)
		return
	}

	err = tmpl.Execute(fil, "")
	if err != nil {
		fmt.Printf("Execution error: %v\n", err)
		return
	}
	fil.Close()

}
