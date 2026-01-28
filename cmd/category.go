package cmd

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/srynprjl/sandwich/internal/category"
)

var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "Manage your categories",
	Long:  `idk`,
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a category",
	Run: func(cmd *cobra.Command, args []string) {
		name, nameErr := cmd.Flags().GetString("name")
		shorthand, shortHandErr := cmd.Flags().GetString("shorthand")
		if nameErr != nil || shortHandErr != nil {
			fmt.Println("Error: Something went wrong!")
			os.Exit(1)
		}
		if name == "" || shorthand == "" {
			fmt.Println("Error: You need to use both --name and --shorthand flag ")
			os.Exit(1)
		}
		c := category.Category{Title: name, Shorthand: shorthand}
		res := c.Add()

		if res["status"] != "201" {
			fmt.Printf("Error: %s", res["message"])
			os.Exit(1)
		}
		fmt.Printf("Success: %s", res["message"])
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a category",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Error: the argument couldn't be converted to string")
			os.Exit(1)
		}
		c := category.Category{Id: id}
		res := c.Delete()
		if res["status"] == "200" {
			fmt.Printf("Success: %s", res["message"])
		} else {
			fmt.Printf("Error: %s", res["message"])
		}
	},
}

var updateCmd = &cobra.Command{
	Use:   "update [id]",
	Short: "Update a category",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := strconv.Atoi(args[0])
		name, nameErr := cmd.Flags().GetString("name")
		shorthand, shortHandErr := cmd.Flags().GetString("shorthand")
		if nameErr != nil || shortHandErr != nil {
			fmt.Println("Error: Something went wrong!")
			os.Exit(1)
		}
		c := category.Category{Id: id, Title: name, Shorthand: shorthand}
		res := c.Update()
		if res["status"] != "200" {
			fmt.Printf("Error: %s\n", res["message"])
			os.Exit(1)
		}
		fmt.Printf("Success: %s\n", res["message"])
	},
}

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View all categories",
	Run: func(cmd *cobra.Command, args []string) {
		res := category.GetAll()
		if res["status"] != "200" {
			fmt.Printf("Error: %s", res["message"])
			os.Exit(1)
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
		fmt.Fprintln(w, "ID\tName\tShorthand")
		for _, data := range res["data"].([]category.Category) {
			fmt.Fprintf(w, "%d\t%s\t%s\n", data.Id, data.Title, data.Shorthand)
		}
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(categoryCmd)
	categoryCmd.AddCommand(addCmd, deleteCmd, updateCmd, viewCmd)

	addCmd.Flags().StringP("name", "n", "", "The name of your category")
	addCmd.Flags().StringP("shorthand", "s", "", "The shorthand name for your category")

	// update
	updateCmd.Flags().StringP("name", "n", "", "The name of your category")
	updateCmd.Flags().StringP("shorthand", "s", "", "The shorthand name for your category")

}
