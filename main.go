/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/srynprjl/sandwich/cmd"
	"github.com/srynprjl/sandwich/utils"
	"github.com/srynprjl/sandwich/utils/db"
)

// import "github.com/srynprjl/sandwich/cmd"

func main() {
	utils.NewConfig()
	utils.InitializeConfig()
	fmt.Println(utils.Conf.Tables["projects"].Default)
	db.InitializeDatabase()
	cmd.Execute()
}
