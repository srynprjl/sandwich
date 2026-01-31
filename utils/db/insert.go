package db

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func (d *Database) InsertOne(tableName string, values map[string]any) error {
	fields := getFields(tableName)
	var insertQuery []any
	var insertValues []string
	for _, field := range fields {
		var val any
		if field == "uuid" {
			val = uuid.New().String()
		} else {
			v, ok := values[field]
			if !ok {
				data, err := GetDefaultValues(tableName, field)
				if err != nil {
					return err
				}
				val = data
			} else {
				val = v
			}
		}
		insertQuery = append(insertQuery, val)
		insertValues = append(insertValues, "?")
	}

	query := fmt.Sprintf("INSERT into %s(%s) VALUES(%s);", tableName, strings.Join(fields, ","), strings.Join(insertValues, ","))
	err := execute(d, query, insertQuery...)
	if err != nil {
		return err
	}
	return nil
}

func (d Database) InsertMany(name string, values []map[string]any) error {
	for _, data := range values {
		err := d.InsertOne(name, data)
		if err != nil {
			return err
		}
	}
	return nil
}
