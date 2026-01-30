package db

import (
	"errors"
	"fmt"
	"github.com/srynprjl/sandwich/utils"
)

func (db *Database) CreateTable(name string, columns []string, coltype []string, primary_key string, autoincrement []string, unique []string, notNull []string, foreignKey []utils.ForeignKey, defaults []map[string]any) error {
	if len(columns) != len(coltype) {
		return errors.New("difference in number between columns and columns type")
	}
	sql_query := BuildSQLTableQuery(columns, coltype, primary_key, autoincrement, unique, notNull, defaults, foreignKey)
	final_query := fmt.Sprintf("CREATE TABLE %s(%s)", name, sql_query)
	DB.Connect()
	conn := DB.Conn
	_, err := conn.Exec(final_query)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) CreateInitialTables() error {
	for key, data := range utils.Conf.Tables {
		err := DB.CreateTable(key, data.Columns, data.ColumnTypes, data.PrimaryKey, data.AutoIncrement, data.Unique, data.NotNull, data.ForeignKey, data.Default)
		if err != nil {
			return err
		}
	}
	return nil
}
