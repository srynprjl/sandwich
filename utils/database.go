package utils

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Name     string
	Location string
	Conn     *sql.DB
}

var DB = Database{
	Name:     DATABASE_NAME,
	Location: DATABASE_LOCATION,
}

func (d *Database) Connect() {
	_, err := os.Stat(d.Location)
	if os.IsNotExist(err) {
		os.MkdirAll(d.Location, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}
	fullPath := d.Location + d.Name + ".db"
	_, fileErr := os.Stat(fullPath)
	if os.IsNotExist(fileErr) {
		_, err = os.Create(fullPath)
		if err != nil {
			log.Fatal(err)
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

func (d *Database) CreateInitialTables() {
	if d.Conn == nil {
		d.Connect()
	}
	sqlQuery := []string{"CREATE TABLE categories(id INTEGER PRIMARY KEY AUTOINCREMENT, name VARCHAR(50) NOT NULL, shorthand VARCHAR(20) UNIQUE)", "CREATE TABLE projects(id INTEGER PRIMARY KEY AUTOINCREMENT, name VARCHAR(255), description VARCHAR(255), completed BOOLEAN  CHECK (completed IN (0, 1)) , favorite BOOLEAN  CHECK (favorite IN (0, 1)) , path VARCHAR(1000), category INTEGER, FOREIGN KEY (id) REFERENCES categories(id))"}
	for _, sql := range sqlQuery {
		_, err := d.Conn.Exec(sql)
		if err != nil {
			log.Fatal(err)
		}
	}
}
