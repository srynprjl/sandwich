package db

import (
	"errors"
	"fmt"
	"strings"

	"github.com/srynprjl/sandwich/utils/config"
)

func BuildForeignKeyStatement(foreignKey []config.ForeignKey) string {
	var foreign_keys []string
	for _, data := range foreignKey {
		var sb strings.Builder
		fmt.Fprintf(&sb, "FOREIGN KEY (%v) REFERENCES %v (%v) ", data.Field, data.To.Table, data.To.Field)
		if data.OnDelete != "" {
			fmt.Fprintf(&sb, "ON DELETE %v", data.OnDelete)
		}
		foreign_keys = append(foreign_keys, sb.String())
	}
	return strings.Join(foreign_keys, ",")
}

func BuildSQLTableQuery(columns []string, coltype []string, constraints config.Constraints) string {
	var queryData []string
	auto := make(map[string]bool)
	for _, v := range constraints.AutoIncrement {
		auto[v] = true
	}
	uniq := make(map[string]bool)
	for _, v := range constraints.Unique {
		uniq[v] = true
	}
	nnull := make(map[string]bool)
	for _, v := range constraints.NotNull {
		nnull[v] = true
	}

	for i, data := range columns {
		var datatype string
		switch coltype[i] {
		case "string":
			datatype = "VARCHAR(255)"
		case "int":
			datatype = "INTEGER"
		default:
			datatype = strings.ToUpper(coltype[i])
		}
		var sb strings.Builder
		sb.WriteString(data + " " + datatype)
		if constraints.PrimaryKey == data {
			sb.WriteString(" PRIMARY KEY")
		}

		if auto[data] {
			sb.WriteString(" AUTOINCREMENT")
		}
		if uniq[data] {
			sb.WriteString(" UNIQUE")
		}
		if nnull[data] {
			sb.WriteString(" NOT NULL")
		}

		// fix this later... make it more efficient later ig
		if len(constraints.Default) != 0 {
			for i := range constraints.Default {
				if val, ok := constraints.Default[i][data]; ok {
					fmt.Fprintf(&sb, " DEFAULT %v", val)
				}
			}
		}

		queryData = append(queryData, sb.String())
	}
	if len(constraints.ForeignKey) != 0 {
		queryData = append(queryData, BuildForeignKeyStatement(constraints.ForeignKey))
	}
	return strings.Join(queryData, ",")
}

func execute(d *Database, query string, args ...any) error {
	d.Connect()
	conn := d.Conn
	defer d.Close()
	_, err := conn.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func query(d *Database, query string, args ...any) ([]map[string]any, error) {
	d.Connect()
	defer d.Close()
	rows, err := d.Conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var data []map[string]any
	for rows.Next() {
		columns := make([]any, len(cols))
		columnPointers := make([]any, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		rowMap := make(map[string]any)
		for i, colName := range cols {
			val := columns[i]
			rowMap[colName] = val
		}
		data = append(data, rowMap)
	}
	return data, nil

}

func getFields(tableName string) []string {
	return config.DefaultTables[tableName].Columns[1:]
}

func getDefaultValues(tableName string, fieldName string) (any, error) {
	value, ok := config.DefaultTables[tableName].Defaults[fieldName]
	if !ok {
		return "", errors.New(fieldName + " is required.")
	}
	return value, nil
}

func joinStatements(fields map[string]any, join string) (string, []any) {
	var keys []string
	var values []any
	for k, v := range fields {
		keys = append(keys, fmt.Sprintf("%s=?", k))
		values = append(values, v)
	}
	keysStatement := strings.Join(keys, join)
	return keysStatement, values
}
