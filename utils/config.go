package utils

import (
	"errors"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

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
	Columns       []string     `yaml:"columns"`
	ColumnTypes   []string     `yaml:"col_types"`
	PrimaryKey    string       `yaml:"primary_key"`
	AutoIncrement []string     `yaml:"auto_increment"`
	ForeignKey    []ForeignKey `yaml:"foreign_key"`
}

type ForeignKey struct {
	Field string    `yaml:"field"`
	To    Reference `yaml:"to"`
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

func ConfigExists(Location string) bool {
	_, err := os.Stat(Location + "config.yml")
	if os.IsExist(err) {
		return true
	}
	return false
}

func ConfigDirExists(Location string) bool {
	_, err := os.Stat(Location)

	if os.IsExist(err) {
		return true
	}
	return false
}

func CreateConfigDir(Location string) error {
	if !ConfigExists(Location) {
		os.MkdirAll(Location, 0777)
		return nil
	}
	return errors.New("Couldn't create config directory")
}

func NewConfig() {
	if !ConfigExists(CONFIG_LOCATION) {
		if !ConfigDirExists(CONFIG_LOCATION) {
			CreateConfigDir(CONFIG_LOCATION)
		}
		_, err := os.Create(CONFIG_LOCATION + "config.yml")
		if err != nil {
			log.Fatal(err)
		}
		DefaultConfig := Config{
			Version:         "0.0.1",
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
			Tables: map[string]Table{
				"categories": Table{
					Columns:       []string{"uuid", "id", "name", "shorthand", "description"},
					ColumnTypes:   []string{"string", "int", "string", "string", "string"},
					PrimaryKey:    "uuid",
					AutoIncrement: []string{"id"},
				},
				"projects": Table{
					Columns:       []string{"uuid", "id", "name", "shorthand", "description", "path", "favorite", "progress", "released", "github", "url", "category"},
					ColumnTypes:   []string{"string", "int", "string", "string", "string", "string", "boolean", "boolean", "boolean", "string", "string", "string"},
					PrimaryKey:    "uuid",
					AutoIncrement: []string{"id"},
					ForeignKey:    []ForeignKey{ForeignKey{Field: "category", To: Reference{Table: "categories", Field: "uuid"}}},
				},
			},
		}
		data, yamlErr := yaml.Marshal(DefaultConfig)
		if yamlErr != nil {
			log.Fatal(yamlErr)
		}
		os.WriteFile(CONFIG_LOCATION+"config.yml", data, 0)
	}

}

func LoadConfig(Location string) (Config, error) {
	c := Config{}
	yf, err := os.ReadFile(Location + "config.yml")
	if err != nil {
		return c, err
	}
	yamlErr := yaml.Unmarshal(yf, &c)
	if yamlErr != nil {
		return c, err
	}
	return c, nil
}
