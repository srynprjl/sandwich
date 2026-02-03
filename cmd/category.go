package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/srynprjl/sandwich/internal/category"
	"github.com/srynprjl/sandwich/utils/config"
	"github.com/srynprjl/sandwich/utils/db"
)

var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "Manage your categories",
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a category",
	Run: func(cmd *cobra.Command, args []string) {
		var data = make(map[string]any)
		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			if f.Shorthand != "h" {
				data[f.Name] = f.Value
			}
		})
		var c category.Category
		res := c.Add(data)
		if res["status"] == "201" {
			fmt.Println("Success: " + res["message"].(string))
			return
		}
		fmt.Println("Failed: " + res["message"].(string))
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete [id | uid]",
	Short: "Delete a category",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := getCategoryForCondition(args)
		res := c.Delete()
		if res["status"] != "200" {
			fmt.Println("Failed: " + res["message"].(string))
			return
		}
		fmt.Println("Success: " + res["message"].(string))
	},
}

var updateCmd = &cobra.Command{
	Use:   "update [id | uid]",
	Short: "Update a category",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := getCategoryForCondition(args)
		updateData := make(map[string]any)
		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			if f.Changed {
				updateData[f.Name] = f.Value
			}
		})
		res := c.Update(updateData)
		if res["status"] != "200" {
			fmt.Println("Failed: " + res["message"].(string))
			return
		}
		fmt.Println("Success: " + res["message"].(string))
	},
}

var viewCmd = &cobra.Command{
	Use:   "list",
	Short: "List all categories",
	Run: func(cmd *cobra.Command, args []string) {
		res := category.GetAll()
		if res["status"] != "200" {
			fmt.Println("Failed: " + res["message"].(string))
			return
		}
		t := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintf(t, "ID\tUID\tName\tDescription\n")
		for _, data := range res["data"].([]map[string]any) {
			fmt.Fprintf(t, "%v\t%v\t%v\t%v\n", data["id"], data["shorthand"], data["name"], data["description"])
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
