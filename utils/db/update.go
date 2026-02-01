package db

import (
	"errors"
	"fmt"
)

// todo (as per the project needs)
// func UpdateTable() {

// }

func (d *Database) UpdateItems(tableName string, fields map[string]any, conditions map[string]any) error {
	if len(fields) == 0 || len(conditions) == 0 {
		return errors.New("provide the fields or conditon")
	}
	fkeysStatement, fvalues := joinStatements(fields, ",")
	ckeysStatement, cvalues := joinStatements(conditions, " and ")
	sqlStatement := fmt.Sprintf("UPDATE %s SET %s WHERE %s", tableName, fkeysStatement, ckeysStatement)
	values := append(fvalues, cvalues...)
	return execute(d, sqlStatement, values...)
}
