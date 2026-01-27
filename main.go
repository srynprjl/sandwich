package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/srynprjl/sandwich/utils"
)

func main() {
	db := utils.Database{
		Name:     utils.DATABASE_NAME,
		Location: utils.DATABASE_LOCATION,
	}
	db.Connect()
	db.CreateInitialTables()
}
