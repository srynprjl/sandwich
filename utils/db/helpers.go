package db

import (
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
