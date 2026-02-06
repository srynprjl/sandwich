package main

import (
	"github.com/srynprjl/stack/cmd"
	"github.com/srynprjl/stack/internal/config"
	"github.com/srynprjl/stack/internal/utils/db"
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
