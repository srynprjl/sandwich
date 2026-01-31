package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/srynprjl/sandwich/internal/category"
	"github.com/srynprjl/sandwich/internal/projects"
	"github.com/srynprjl/sandwich/utils/config"
)

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage your projects with ease",
	Run: func(cmd *cobra.Command, args []string) {
		fav, favErr := cmd.Flags().GetBool("favorite")
		comp, comErr := cmd.Flags().GetBool("completed")
		projectMap := make(map[string]bool)
		if favErr != nil {
			fmt.Printf("Error: %s", favErr.Error())
		}
		if comErr != nil {
			fmt.Printf("Error: %s", comErr.Error())
		}
		if !fav && !comp {
			cmd.Help()
			return
		}
		projectMap["favorite"] = fav
		projectMap["completed"] = comp
		res := projects.GetProjectWhere(projectMap)
		if res["status"] != "200" {
			fmt.Printf("Error: %s", res["message"])
			os.Exit(1)
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(w, "ID\tName\tPath\tCompleted\tFavourite\tDescription")
		for _, data := range res["data"].([]projects.Project) {
			fmt.Fprintf(w, "%d\t%s\t%t\t%t\t%s\n", data.Id, data.Title, data.Completed, data.Favourite, data.Description)
		}
		w.Flush()
	},
}

var projectAddCmd = &cobra.Command{
	Use:   "add [categoryId]",
	Short: "Add a project",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		catId, catErr := strconv.Atoi(args[0])
		if catErr != nil {
			fmt.Println("Error: couldn't convert id to integer")
			os.Exit(1)
		}
		p := projects.Project{Category: catId}
		newData := make(map[string]any)
		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			newData[f.Name] = f.Value
		})
		byteData, _ := json.Marshal(newData)
		err := json.Unmarshal(byteData, &p)
		if err != nil {
			fmt.Println("Error: couldn't convert data to object")
			os.Exit(1)
		}
		res := p.Add()
		if res["status"] != "201" {
			fmt.Printf("Error: %s", res["message"])
			os.Exit(1)
		}
		fmt.Printf("Success: %s", res["message"])
	},
}

var projectDeleteCmd = &cobra.Command{
	Use:   "delete [categoryId] [id]",
	Short: "Delete the project",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[1])
		catId, catErr := strconv.Atoi(args[0])
		if err != nil || catErr != nil {
			fmt.Println("Error: couldn't convert id to integer")
			os.Exit(1)
		}
		p := projects.Project{Id: id, Category: catId}
		res := p.Remove()
		if res["status"] != "200" {
			fmt.Printf("Error: %s", res["message"])
			os.Exit(1)
		}
		fmt.Printf("Success: %s", res["message"])
	},
}

var projectUpdateCmd = &cobra.Command{
	Use:   "update [categoryId] [id]",
	Short: "Update the information about the project",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[1])
		catId, catErr := strconv.Atoi(args[0])
		if err != nil || catErr != nil {
			fmt.Println("Error: couldn't convert id to integer")
			os.Exit(1)
		}
		p := projects.Project{Id: id, Category: catId}
		newData := make(map[string]any)
		cmd.Flags().Visit(func(f *pflag.Flag) {
			newData[f.Name] = f.Value
		})
		res := p.Update(newData)
		if res["status"] != "200" {
			fmt.Printf("Error: %s\n", res["message"])
			os.Exit(1)
		}
		fmt.Printf("Success: %s\n", res["message"])
	},
}

var projectViewCmd = &cobra.Command{
	Use:   "view [categoryId] [id]",
	Short: "View information about the project",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[1])
		catId, catErr := strconv.Atoi(args[0])
		if err != nil || catErr != nil {
			fmt.Println("Error: couldn't convert id to integer")
			os.Exit(1)
		}
		p := projects.Project{Id: id, Category: catId}
		res := p.Get()
		if res["status"] != "200" {
			fmt.Printf("Error: %s\n", res["message"])
			os.Exit(1)
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "Column\tValue")
		t := reflect.TypeOf(res["data"].(projects.Project))
		v := reflect.ValueOf(res["data"].(projects.Project))
		for i := range t.NumField() {
			fieldType := t.Field(i)
			valueType := v.Field(i)
			fmt.Fprintf(w, "%s\t%v\n", fieldType.Name, valueType.Interface())
		}
		w.Flush()
	},
}

var projectListAllCmd = &cobra.Command{
	Use:   "list [category]",
	Short: "List all the project in the category",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		categoryId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: couldn't convert id to integer")
			os.Exit(1)
		}
		res := projects.GetProjects(category.Category{Id: categoryId})
		if res["status"] != "200" {
			fmt.Printf("Error: %s\n", res["message"])
			os.Exit(1)
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(w, "ID\tName\tPath\tCompleted\tFavourite\tDescription")
		for _, data := range res["data"].([]projects.Project) {
			fmt.Fprintf(w, "%d\t%s\t%t\t%t\t%s\n", data.Id, data.Title, data.Completed, data.Favourite, data.Description)
		}
		w.Flush()

	},
}

var projectEditCmd = &cobra.Command{
	Use:   "edit [id]",
	Short: "Edit the project using your default editor",
	Run: func(cmd *cobra.Command, args []string) {
		defaultEditor := os.Getenv("EDITOR")
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: couldn't convert id to integer")
			os.Exit(1)
		}
		p := projects.Project{Id: id}
		res := p.GetField([]string{"path"})
		if res["status"] != "200" {
			fmt.Printf("Error: %s\n", res["message"])
			os.Exit(1)
		}
		path := res["data"].([]any)[0].(string)
		fmt.Print(defaultEditor)
		if defaultEditor == "" {
			fmt.Println("Error: No default editor set. ")
			os.Exit(1)
		}
		editor := exec.Command(defaultEditor, path)

		editor.Stdin = os.Stdin
		editor.Stdout = os.Stdout
		editor.Stderr = os.Stderr

		editErr := editor.Run()
		if editErr != nil {
			fmt.Println(editErr.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(projectCmd)
	projectCmd.AddCommand(projectAddCmd, projectDeleteCmd, projectEditCmd, projectUpdateCmd, projectViewCmd, projectListAllCmd)
	projectCmd.Flags().BoolP("favorite", "f", false, "List all favorite projects")
	projectCmd.Flags().BoolP("completed", "c", false, "List all completed projects")

	projectAddCmd.Flags().StringP("name", "n", "", "The name of the project")
	projectAddCmd.Flags().StringP("description", "d", "", "The description of the project")
	projectAddCmd.Flags().StringP("path", "p", config.Conf.ProjectLocation, "The path of the project")
	projectAddCmd.Flags().BoolP("favorite", "f", false, "Is the project your favourite.")
	projectAddCmd.Flags().BoolP("completed", "c", false, "Is the project finished.")

	projectUpdateCmd.Flags().StringP("name", "n", "", "The name of the project")
	projectUpdateCmd.Flags().StringP("description", "d", "", "The description of the project")
	projectUpdateCmd.Flags().StringP("path", "p", "", "The path of the project")
	projectUpdateCmd.Flags().BoolP("favorite", "f", false, "Is the project your favourite.")
	projectUpdateCmd.Flags().BoolP("completed", "o", false, "Is the project finished.")
	projectUpdateCmd.Flags().IntP("category", "c", 0, "Which category does it fall in?")
}
