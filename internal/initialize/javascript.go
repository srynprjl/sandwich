package initialize

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func installDependencies(toolkit string) {
	err := runCommand(toolkit, "i")
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
}

func checkNPMToolkits() (string, error) {
	toolkits := []string{"bun", "pnpm", "npm", "yarn"}
	for _, toolkit := range toolkits {
		exists := checkDependencies(toolkit)
		if exists {
			return toolkit, nil
		}
	}
	return "", errors.New("no toolkit found")
}

func checkFramework(f string) (string, error) {
	switch f {
	case "react", "vanilla", "vue", "preact", "lit", "svelte", "solid", "qwik", "next":
		return f, nil
	default:
		return "", errors.New("invalid framework")
	}

}

func InitJS(project_data map[string]any, framework string, lang string) {

	initiated := checkProjectInitiated(project_data["path"].(string))
	if initiated {
		fmt.Println("The project files already exist!!! ")
		return
	}
	// check if toolkit exists
	toolkit, err := checkNPMToolkits()
	if err != nil {
		fmt.Println("Error: No toolkit found. Please install one of the following [bun, pnpm, yarn, npm]")
		return
	}
	framework, err = checkFramework(framework)
	if err != nil {
		fmt.Println("error: Please input a correct framework")
		return
	}

	switch framework {
	case "next", "nextjs":
		InitNext(toolkit, lang, project_data)
	default:
		InitViteToolkits(project_data, toolkit, lang, framework)
	}

	fmt.Println("Initialized. ")

}

func InitViteToolkits(project_data map[string]any, toolkit string, lang string, framework string) {
	templates := framework
	if lang == "ts" || lang == "typescript" {
		templates += "-ts"
	}
	os.Chdir(project_data["path"].(string))
	err := runCommand(toolkit, "create", "vite@latest", "--template", templates, "--rolldown", "--no-interactive", ".")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	installDependencies(toolkit)
}

func prepareName(name string) string {
	data := strings.ToLower(name)
	data = strings.ReplaceAll(data, " ", "")
	return data
}

func InitNext(toolkit string, lang string, project_data map[string]any) {
	name := prepareName(project_data["name"].(string))
	args := []string{"create", "next-app@latest", name, "--react-compiler", "--tailwind", "--biome", "--app", "--src-dir", "--turbopack", "--yes"}
	if lang == "js" || lang == "javascript" {
		args = append(args, "--js")
	} else {
		args = append(args, "--ts")
	}
	args = append(args, "--use-"+toolkit)
	projectDir := project_data["path"].(string)
	parentDir := filepath.Dir(filepath.Clean(projectDir))
	err := os.Chdir(parentDir)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	err = runCommand(toolkit, args...)

	os.Remove(projectDir)
	os.Rename(path.Join(parentDir, name), project_data["name"].(string))
	os.Chdir(projectDir)

	installDependencies(toolkit)
}
