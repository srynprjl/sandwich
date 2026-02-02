package db

import (
	"errors"
	"fmt"

	"github.com/srynprjl/sandwich/utils/config"
)

func (db *Database) CreateTable(name string, columns []string, coltype []string, constraints config.Constraints) error {
	if len(columns) != len(coltype) {
		return errors.New("difference in number between columns and columns type")
	}
	final_query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(%s)", name, BuildSQLTableQuery(columns, coltype, constraints))
	err := execute(db, final_query)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) CreateInitialTables() error {
	for key, data := range config.DefaultTables {
		err := DB.CreateTable(key, data.Columns, data.ColumnTypes, data.Constraints)
		if err != nil {
			return err
		}
	}
	return nil
}
