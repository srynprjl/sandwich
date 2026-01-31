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
	Columns     []string    `yaml:"columns"`
	ColumnTypes []string    `yaml:"col_types"`
	Constraints Constraints `yaml:"constraints"`
}

type Constraints struct {
	PrimaryKey    string           `yaml:"primary_key"`
	AutoIncrement []string         `yaml:"auto_increment"`
	ForeignKey    []ForeignKey     `yaml:"foreign_key"`
	Unique        []string         `yaml:"unique"`
	NotNull       []string         `yaml:"not_null"`
	Default       []map[string]any `yaml:"default"`
}

type ForeignKey struct {
	Field    string    `yaml:"field"`
	To       Reference `yaml:"to"`
	OnDelete string    `yaml:"on_delete"`
}

type Reference struct {
	Table string `yaml:"table"`
	Field string `yaml:"field"`
}

type Config struct {
	Version         string           `yaml:"version"`
	ProjectLocation string           `yaml:"project_location"`
	Database        DatabaseConfig   `yaml:"db"`
	Web             WebConfig        `yaml:"web"`
	Tables          map[string]Table `yaml:"tables"`
}
