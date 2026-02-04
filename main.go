package main

import (
	"github.com/srynprjl/sandwich/cmd"
	"github.com/srynprjl/sandwich/internal/config"
	"github.com/srynprjl/sandwich/internal/utils/db"
)

func startup() {
	config.NewConfig()
	config.InitializeConfig()
	db.InitializeDatabase()
}

func main() {
	startup()
	cmd.Execute()
}
