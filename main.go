/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/srynprjl/sandwich/cmd"
	"github.com/srynprjl/sandwich/utils"
)

// import "github.com/srynprjl/sandwich/cmd"

func main() {
	utils.NewConfig()
	utils.InitializeVariables()
	cmd.Execute()
}
