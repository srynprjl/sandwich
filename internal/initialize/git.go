package initialize

import (
	"fmt"
	"os"
	"os/exec"
)

func isGitInitialized(project_path string) bool {
	_, err := os.Stat(project_path + ".git/")
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func InitializeGitRepo(project_path string) {
	initialized := isGitInitialized(project_path)
	if !initialized {
		err := os.Chdir(project_path)
		if err != nil {
			fmt.Println("Error: ", err.Error())
			return
		}
		gitInit := exec.Command("git", "init")
		gitErr := gitInit.Run()
		if gitErr != nil {
			fmt.Println("Error: ", gitErr.Error())
			return
		}
	}
}

func SetGitOrigin(data map[string]any) {
	if data["github"] != "" {
		// git set remote origin
	}
}

func DeleteGitRepo() {

}

// func GetCommitHistory() {

// }

// func GetIssues() {

// }
