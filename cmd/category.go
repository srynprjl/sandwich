package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/srynprjl/sandwich/internal/category"
	"github.com/srynprjl/sandwich/internal/config"
	"github.com/srynprjl/sandwich/internal/utils/db"
)

var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "Manage your categories",
}

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"push"},
	Short:   "Add a category",
	Run: func(cmd *cobra.Command, args []string) {
		var data = make(map[string]any)
		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			if f.Shorthand != "h" {
				data[f.Name] = f.Value
			}
		})
		var c category.Category
		res := c.Add(data)
		if res.Error != nil {
			fmt.Println("Failed: " + res.Message)
			return
		}
		fmt.Println("Success: " + res.Message)

	},
}
var deleteCmd = &cobra.Command{
	Use:     "delete [uid]",
	Aliases: []string{"pop"},
	Short:   "Delete a category",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := category.Category{Shorthand: args[0]}
		res := c.Delete()
		if res.Error != nil {
			fmt.Println("Failed: " + res.Message)
			return
		}
		fmt.Println("Success: " + res.Message)
	},
}

var updateCmd = &cobra.Command{
	Use:     "update [uid]",
	Aliases: []string{"patch"},
	Short:   "Update a category",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := category.Category{Shorthand: args[0]}
		updateData := make(map[string]any)
		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			if f.Changed {
				updateData[f.Name] = f.Value
			}
		})
		res := c.Update(updateData)
		if res.Error != nil {
			fmt.Println("Failed: " + res.Message)
			return
		}
		fmt.Println("Success: " + res.Message)
	},
}

var viewCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"trace"},
	Short:   "List all categories",
	Run: func(cmd *cobra.Command, args []string) {
		data, res := category.GetAll()
		if res.Error != nil {
			fmt.Println("Failed: " + res.Message)
			return
		}
		t := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintf(t, "Index\tUID\tName\tDescription\n")
		for i, data := range data {
			fmt.Fprintf(t, "%v\t%v\t%v\t%v\n", i+1, data["shorthand"], data["name"], data["description"])
		}
		t.Flush()

	},
}

func init() {
	rootCmd.AddCommand(categoryCmd)
	categoryCmd.AddCommand(addCmd, deleteCmd, updateCmd, viewCmd)
	fields := config.DefaultTables["categories"].Columns[2:]
	for _, data := range fields {
		var def, ok = db.GetDefaultValues("categories", data)
		if ok == nil {
			switch def := def.(type) {
			case string:
				addCmd.Flags().String(data, def, fmt.Sprintf("The %s of your category", data))
				updateCmd.Flags().String(data, "", fmt.Sprintf("The %s of your category", data))
			case int:
				addCmd.Flags().Int(data, def, fmt.Sprintf("The %s of your category", data))
				updateCmd.Flags().Int(data, 0, fmt.Sprintf("The %s of your category", data))
			case bool:
				addCmd.Flags().Bool(data, def, fmt.Sprintf("The %s of your category", data))
				updateCmd.Flags().Bool(data, false, fmt.Sprintf("The %s of your category", data))
			}
		} else {
			addCmd.Flags().String(data, "", fmt.Sprintf("The %s of your category", data))
			addCmd.MarkFlagRequired(data)
			updateCmd.Flags().String(data, "", fmt.Sprintf("The %s of your category", data))
		}
	}

}
