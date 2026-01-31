package config

var DefaultConfig = Config{
	ProjectLocation: "/home/srynprjl/.local/development/projects/",
	Database: DatabaseConfig{
		Type:     "sqlite",
		Location: "/home/srynprjl/.local/share/sandwich/",
		Name:     "sandwich_test",
	},
	Web: WebConfig{
		Host: "127.0.0.1",
		Port: 5000,
	},
}

var DefaultTables = map[string]Table{
	"categories": Table{
		Columns:     []string{"id", "uuid", "name", "shorthand", "description"},
		ColumnTypes: []string{"int", "string", "string", "string", "string"},
		Constraints: Constraints{
			PrimaryKey:    "id",
			AutoIncrement: []string{"id"},
			Unique:        []string{"uuid"},
		},
	},
	"projects": Table{
		Columns:     []string{"id", "uuid", "name", "shorthand", "description", "path", "favorite", "progress", "released", "github", "url", "category"},
		ColumnTypes: []string{"int", "string", "string", "string", "string", "string", "boolean", "boolean", "boolean", "string", "string", "int"},
		Constraints: Constraints{
			PrimaryKey:    "id",
			AutoIncrement: []string{"id"},
			Unique:        []string{"uuid"},
			ForeignKey:    []ForeignKey{{Field: "category", To: Reference{Table: "categories", Field: "id"}}},
			Default:       []map[string]any{{"category": 0}},
		},
	},
}
