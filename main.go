/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/srynprjl/sandwich/cmd"
	"github.com/srynprjl/sandwich/utils/config"
	"github.com/srynprjl/sandwich/utils/db"
)

// import "github.com/srynprjl/sandwich/cmd"
func startup() {
	config.NewConfig()
	config.InitializeConfig()
	db.InitializeDatabase()
}

func main() {
	startup()
	cmd.Execute()
}
