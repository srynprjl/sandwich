package db

import (
	"fmt"
	"strings"
)

func (d *Database) Query(tableName string, fields []string, condition map[string]any, extras ...string) ([]map[string]any, error) {
	var field string
	var statement string
	var values []any
	if len(fields) == 0 {
		field = "*"
	} else {
		field = strings.Join(fields, ",")
	}
	sqlStatement := fmt.Sprintf("SELECT %s FROM %s", field, tableName)
	if len(condition) != 0 {
		statement, values = joinStatements(condition, " and ")
		sqlStatement = fmt.Sprintf("%s WHERE %s", sqlStatement, statement)
	}
	if len(extras) != 0 {
		sqlStatement = fmt.Sprintf("%s %s", sqlStatement, strings.Join(extras, " "))
	}
	return query(d, sqlStatement, values...)
}

func (d *Database) QueryLimit(tableName string, fields []string, condition map[string]any, num int) ([]map[string]any, error) {
	return d.Query(tableName, fields, condition, fmt.Sprintf("LIMIT %d", num))
}

func (d *Database) QueryRandom(tableName string, fields []string, condition map[string]any, num int) ([]map[string]any, error) {
	return d.Query(tableName, fields, condition, "ORDER BY RANDOM()", fmt.Sprintf("LIMIT %d", num))
}

func (d *Database) CheckExists(tableName string, data map[string]any) (bool, error) {
	dat, err := d.QueryLimit(tableName, []string{"1"}, data, 1)
	if err != nil {
		return false, err
	}
	if len(dat) == 0 {
		return false, nil
	}
	return true, nil
}

func (d *Database) CheckTableExists(tableName string) (bool, error) {
	dat, err := query(d, "SELECT name FROM sqlite_master WHERE type='table' AND name=?", tableName)
	if err != nil {
		return false, err
	}
	if len(dat) == 0 {
		return false, nil
	}
	return true, nil
}
