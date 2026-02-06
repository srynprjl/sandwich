package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/srynprjl/stack/internal/utils/exports"
)

var rootCmd = &cobra.Command{
	Use:   "stack",
	Short: "A project management app for personal use",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			version = "alpha"
		)
		if v, _ := cmd.Flags().GetBool("version"); v {
			fmt.Printf("stack %v\nCopyright (C) 2026 sysnefo.", version)
		} else {
			cmd.Help()
		}

	},
}

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import all data from a JSON file",
	Run: func(cmd *cobra.Command, args []string) {
		fileFormat, ffErr := cmd.Flags().GetString("format")
		if ffErr != nil {
			fmt.Println("Error: " + ffErr.Error())
			return
		}
		path, pErr := cmd.Flags().GetString("path")
		if pErr != nil {
			fmt.Println("Error: " + pErr.Error())
			return
		}
		err := exports.Import(fileFormat, path)
		if err != nil {

			return
		}
	},
}

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export all data into a JSON file",
	Run: func(cmd *cobra.Command, args []string) {
		fileFormat, ffErr := cmd.Flags().GetString("format")
		if ffErr != nil {
			fmt.Println("Error: " + ffErr.Error())
			return
		}
		path, pErr := cmd.Flags().GetString("path")
		if pErr != nil {
			fmt.Println("Error: " + pErr.Error())
			return
		}

		name, nErr := cmd.Flags().GetString("name")
		if nErr != nil {
			fmt.Println("Error: " + nErr.Error())
			return
		}

		tables, tErr := cmd.Flags().GetStringSlice("tables")
		if tErr != nil {
			fmt.Println("Error: " + tErr.Error())
			return
		}

		err := exports.Export(fileFormat, path, name, tables...)
		if err != nil {
			fmt.Println("Error: " + err.Error())
			return
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Version of sandwich")

	importCmd.Flags().StringP("format", "f", "json", "The format of the file you trying to import")
	importCmd.Flags().StringP("path", "p", "", "The file you trying to import from")
	importCmd.MarkFlagRequired("path")

	userDir, err := os.UserHomeDir()
	if err != nil {
		userDir = "./"
	}
	exportCmd.Flags().StringP("format", "f", "json", "which format u trynna export the file")
	exportCmd.Flags().StringP("path", "p", userDir, "where u saving the file")
	exportCmd.Flags().StringP("name", "n", "export", "name of file u trynna export")
	exportCmd.Flags().StringSliceP("tables", "t", []string{"all"}, "which tables you trying to export")
	rootCmd.AddCommand(importCmd, exportCmd)
}
