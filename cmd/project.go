package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/srynprjl/stack/internal/category"
	"github.com/srynprjl/stack/internal/config"
	"github.com/srynprjl/stack/internal/initialize"
	"github.com/srynprjl/stack/internal/projects"
	"github.com/srynprjl/stack/internal/utils/db"
)

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage your projects with ease",
	Run: func(cmd *cobra.Command, args []string) {
		projectMap := make(map[string]any)
		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			if f.Changed {
				projectMap[f.Name] = f.Value
			}
		})
		if len(projectMap) == 0 {
			cmd.Help()
			return
		}
		data, res := projects.GetProjectWhere(projectMap)
		if res.Error != nil {
			fmt.Printf("Error: %s\n", res.Message)
			return
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(w, "Index\tUID\tName\tInProgress\tReleased\tFavourite\tDescription")
		for i, data := range data {
			fmt.Fprintf(w, "%d\t%s\t%s\t%t\t%t\t%t\t%s\n", i+1, data["shorthand"], data["name"], data["progress"], data["released"], data["favorite"], data["description"])
		}
		w.Flush()
	},
}

var projectAddCmd = &cobra.Command{
	Use:     "add [categoryUID]",
	Aliases: []string{"push"},
	Short:   "Add a project",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := category.Category{Shorthand: args[0]}
		data, resp := c.GetField([]string{"id"})
		if resp.Error != nil {
			fmt.Println("Error: " + resp.Message)
		}
		cats := int(data["id"].(int64))
		p := projects.Project{Category: cats}
		newData := make(map[string]any)
		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			newData[f.Name] = f.Value
		})
		newData["category"] = cats
		res := p.Add(newData)
		if res.Error != nil {
			fmt.Printf("Error: %s\n", res.Message)
			return
		}
		fmt.Printf("Success: %s\n", res.Message)
	},
}

var projectDeleteCmd = &cobra.Command{
	Use:     "delete [uid]",
	Aliases: []string{"pop"},
	Short:   "Delete the project",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		p := projects.Project{ProjectId: args[0]}
		res := p.Remove()
		if res.Error != nil {
			fmt.Printf("Error: %s\n", res.Message)
			return
		}
		fmt.Printf("Success: %s\n", res.Message)
	},
}

var projectUpdateCmd = &cobra.Command{
	Use:     "update [uid]",
	Aliases: []string{"patch"},
	Short:   "Update the information about the project",
	Run: func(cmd *cobra.Command, args []string) {
		p := projects.Project{ProjectId: args[0]}
		newData := make(map[string]any)
		cmd.Flags().Visit(func(f *pflag.Flag) {
			if f.Changed {
				newData[f.Name] = f.Value
			}
		})
		res := p.Update(newData)
		if res.Error != nil {
			fmt.Printf("Error: %s\n", res.Message)
			return
		}
		fmt.Printf("Success: %s\n", res.Message)
	},
}

var projectViewCmd = &cobra.Command{
	Use:     "view [uid]",
	Aliases: []string{"peek"},
	Short:   "View information about the project",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		p := projects.Project{ProjectId: args[0]}
		data, res := p.Get()
		if res.Error != nil {
			fmt.Printf("Error: %s\n", res.Message)
			return
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "Column\tValue")
		var project projects.Project
		byteData, _ := json.Marshal(data)
		json.Unmarshal(byteData, &project)
		t := reflect.TypeFor[projects.Project]()
		v := reflect.ValueOf(project)
		for i := range t.NumField() {
			fieldType := t.Field(i)
			valueType := v.Field(i)
			fmt.Fprintf(w, "%s\t%v\n", fieldType.Name, valueType)
		}
		w.Flush()
	},
}

var projectListAllCmd = &cobra.Command{
	Use:     "list [categoryUID]",
	Aliases: []string{"trace"},
	Short:   "List all the project in the category",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := category.Category{Shorthand: args[0]}
		data, res := projects.GetProjects(c)
		if res.Error != nil {
			fmt.Printf("Error: %s\n", res.Message)
			return
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(w, "Index\tUID\tName\tInProgress\tReleased\tFavourite\tDescription")
		for i, data := range data {
			fmt.Fprintf(w, "%d\t%s\t%s\t%t\t%t\t%t\t%s\n", i+1, data["shorthand"], data["name"], data["progress"], data["released"], data["favorite"], data["description"])
		}
		w.Flush()
	},
}

var projectEditCmd = &cobra.Command{
	Use:   "edit [uid]",
	Short: "Edit the project using your default editor",
	Run: func(cmd *cobra.Command, args []string) {
		defaultEditor := os.Getenv("EDITOR")
		p := projects.Project{ProjectId: args[0]}
		data, res := p.GetField([]string{"path"})
		if res.Error != nil {
			fmt.Printf("Error: %s\n", res.Message)
			return
		}
		path := data["path"].(string)
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

var projectInitCommand = &cobra.Command{
	Use:   "init [uid]",
	Short: "Initialize a project",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		p := projects.Project{ProjectId: args[0]}
		var argsCorrect bool = false
		str, _ := cmd.Flags().GetString("lang")
		switch str {
		case
			"go", "golang",
			"rust",
			"java",
			"kotlin",
			"python", "py",
			"js", "javascript":
			argsCorrect = true
		}
		if !argsCorrect {
			fmt.Println("No support for the language.")
			return
		}
		initialize.Init(str, p)
	},
}

func init() {
	rootCmd.AddCommand(projectCmd)
	projectCmd.AddCommand(projectAddCmd, projectDeleteCmd, projectEditCmd, projectUpdateCmd, projectViewCmd, projectListAllCmd, projectInitCommand)

	projectCmd.Flags().BoolP("favorite", "f", false, "List all favorite projects")
	projectCmd.Flags().BoolP("released", "r", false, "List all completed projects")
	projectCmd.Flags().BoolP("progress", "p", false, "List all projects that are in progress")

	projectInitCommand.Flags().StringP("lang", "l", "", "Which language to make?")
	projectInitCommand.MarkFlagRequired("lang")
	fields := config.DefaultTables["projects"].Columns[2:]
	for _, data := range fields {
		var def, ok = db.GetDefaultValues("projects", data)
		if ok == nil {
			switch def := def.(type) {
			case string:
				projectAddCmd.Flags().String(data, def, fmt.Sprintf("The %s of your project", data))
				projectUpdateCmd.Flags().String(data, "", fmt.Sprintf("The %s of your project", data))
			case int:
				if data != "category" {
					projectAddCmd.Flags().Int(data, def, fmt.Sprintf("The %s of your project", data))
				}
				projectUpdateCmd.Flags().Int(data, 0, fmt.Sprintf("The %s of your project", data))
			case bool:
				projectAddCmd.Flags().Bool(data, def, fmt.Sprintf("The %s of your project", data))
				projectUpdateCmd.Flags().Bool(data, false, fmt.Sprintf("The %s of your project", data))
			}
		} else {
			projectAddCmd.Flags().String(data, "", fmt.Sprintf("The %s of your project (Required)", data))
			projectAddCmd.MarkFlagRequired(data)
			projectUpdateCmd.Flags().String(data, "", fmt.Sprintf("The %s of your project", data))
		}
	}
}
