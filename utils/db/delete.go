package db

import (
	"errors"
	"fmt"
)

func (d *Database) DeleteTable(tableName string) error {
	sqlStatement := "DROP TABLE ? "
	err := execute(d, sqlStatement, tableName)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) DeleteItem(tableName string, fields map[string]any) error {
	if len(fields) == 0 {
		return errors.New("condition must be provided")
	}
	keysStatement, values := joinStatements(fields, " and ")
	sqlStatement := fmt.Sprintf("DELETE FROM %s WHERE %s", tableName, keysStatement)
	return execute(d, sqlStatement, values...)
}
