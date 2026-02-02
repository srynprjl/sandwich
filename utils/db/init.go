package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/srynprjl/sandwich/utils/config"
)

type Database struct {
	Name     string
	Location string
	Conn     *sql.DB
}

var DB Database

func InitializeDatabase() {
	DB = Database{
		Name:     config.Conf.Database.Name,
		Location: config.Conf.Database.Location,
	}
}

func (d *Database) Connect() {

	_, err := os.Stat(d.Location)
	if os.IsNotExist(err) {
		os.MkdirAll(d.Location, 0777)
	}
	fullPath := d.Location + d.Name + ".db"
	_, fileErr := os.Stat(fullPath)
	if os.IsNotExist(fileErr) {
		_, err = os.Create(fullPath)
		if err != nil {
			log.Fatal(err)
		}
		dbErr := d.CreateInitialTables()
		if dbErr != nil {
			log.Fatal(dbErr)
		}
	}
	conn, err := sql.Open("sqlite3", fullPath)
	if err != nil {
		log.Fatal(err)
	}
	_, pragmaErr := conn.Exec("PRAGMA foreign_keys = ON")
	d.Conn = conn
	if pragmaErr != nil {
		print("FOREIGN KEY FAILED")
	}

}

func (d *Database) Close() {
	if d.Conn != nil {
		d.Conn.Close()
	}
}
