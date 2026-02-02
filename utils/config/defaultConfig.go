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
	Bot: DiscordBot{
		Token: "",
	},
}

var DefaultTables = map[string]Table{
	"categories": Table{
		Columns:     []string{"id", "uuid", "name", "shorthand", "description"},
		ColumnTypes: []string{"int", "string", "string", "string", "string"},
		Defaults:    map[string]any{"name": "", "description": "This is a new category"},
		Constraints: Constraints{
			PrimaryKey:    "id",
			AutoIncrement: []string{"id"},
			Unique:        []string{"uuid", "shorthand"},
		},
	},
	"projects": Table{
		Columns:     []string{"id", "uuid", "name", "shorthand", "description", "path", "favorite", "progress", "released", "github", "url", "category"},
		ColumnTypes: []string{"int", "string", "string", "string", "string", "string", "boolean", "boolean", "boolean", "string", "string", "int"},
		Defaults:    map[string]any{"name": "New Project", "description": "A new project", "path": Conf.ProjectLocation, "favorite": false, "progress": false, "released": false, "github": "", "url": "", "category": 0},
		Constraints: Constraints{
			PrimaryKey:    "id",
			AutoIncrement: []string{"id"},
			Unique:        []string{"uuid", "shorthand"},
			ForeignKey:    []ForeignKey{{Field: "category", To: Reference{Table: "categories", Field: "id"}, OnDelete: "SET NULL"}},
			Default:       []map[string]any{{"category": 0}},
		},
	},
}
