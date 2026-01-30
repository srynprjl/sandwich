package db

import (
	"fmt"
	"strings"

	"github.com/srynprjl/sandwich/utils"
)

func BuildForeignKeyStatement(foreignKey []utils.ForeignKey) string {
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

func BuildSQLTableQuery(columns []string, coltype []string, primary_key string, autoincrement []string, unique []string, notNull []string, defaults []map[string]any, foreignKey []utils.ForeignKey) string {
	var queryData []string
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
		if primary_key == data {
			sb.WriteString(" PRIMARY KEY")
		}
		auto := make(map[string]bool)
		for _, v := range autoincrement {
			auto[v] = true
		}
		uniq := make(map[string]bool)
		for _, v := range unique {
			uniq[v] = true
		}
		nnull := make(map[string]bool)
		for _, v := range notNull {
			nnull[v] = true
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

		if len(defaults) != 0 {
			for i := range defaults {
				if val, ok := defaults[i][data]; ok {
					fmt.Fprintf(&sb, " DEFAULT %v", val)
				}
			}
		}
		queryData = append(queryData, sb.String())
	}
	if len(foreignKey) != 0 {
		queryData = append(queryData, BuildForeignKeyStatement(foreignKey))
	}
	return strings.Join(queryData, ",")
}
