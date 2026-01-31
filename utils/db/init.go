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
	d.Conn = conn
}

func (d *Database) Close() {
	if d.Conn != nil {
		d.Conn.Close()
	}
}

// func (d *Database) CreateInitialTables() {
// 	if d.Conn == nil {
// 		d.Connect()
// 	}
// 	sqlQuery := []string{"CREATE TABLE categories(id INTEGER PRIMARY KEY AUTOINCREMENT, name VARCHAR(50) NOT NULL, shorthand VARCHAR(20) UNIQUE)", "CREATE TABLE projects(id INTEGER PRIMARY KEY AUTOINCREMENT, name VARCHAR(255), description VARCHAR(255), completed BOOLEAN  CHECK (completed IN (0, 1)) , favorite BOOLEAN  CHECK (favorite IN (0, 1)) , path VARCHAR(1000), category INTEGER, FOREIGN KEY (id) REFERENCES categories(id))"}
// 	for _, sql := range sqlQuery {
// 		_, err := d.Conn.Exec(sql)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }
