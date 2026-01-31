package config

type DatabaseConfig struct {
	Type     string `yaml:"type"`
	Location string `yaml:"location"`
	Name     string `yaml:"name"`
}

type WebConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Table struct {
	Columns     []string
	ColumnTypes []string
	Defaults    map[string]any
	Constraints Constraints
}

type Constraints struct {
	PrimaryKey    string
	AutoIncrement []string
	ForeignKey    []ForeignKey
	Unique        []string
	NotNull       []string
	Default       []map[string]any
}

type ForeignKey struct {
	Field    string
	To       Reference
	OnDelete string
}

type Reference struct {
	Table string
	Field string
}

type Config struct {
	ProjectLocation string         `yaml:"project_location"`
	Database        DatabaseConfig `yaml:"db"`
	Web             WebConfig      `yaml:"web"`
}
