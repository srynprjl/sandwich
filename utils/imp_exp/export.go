package impexp

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/srynprjl/sandwich/internal/logic"
	"github.com/srynprjl/sandwich/utils/db"
	"gopkg.in/yaml.v3"
)

func Export(fileFormat string, path string, fileName string, tables ...string) error {

	// check if file format correct\
	// check if path exists
	_, pathErr := os.Stat(path)
	if os.IsNotExist(pathErr) {
		return errors.New("The path doesn't exist")
	}

	if len(tables) == 0 || (tables[0] == "all" && len(tables) == 1) {
		tables = []string{"categories", "projects"}
	}

	a := []map[string]any{}
	for _, table := range tables {
		ok, err := db.DB.CheckTableExists(table)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("The table doesn't exist")
		}

		// create an instance of the file format
		data, err := db.DB.Query(table, []string{}, map[string]any{})

		if err != nil {
			return err
		}
		var cats []logic.Category
		var proj []logic.Project
		// take data from that table and map to the map or struct idk
		switch fileFormat {
		case "json":
			for _, d := range data {
				mapData, err := json.Marshal(d)
				if err != nil {
					return err
				}
				if table == "categories" {
					var c logic.Category
					json.Unmarshal(mapData, &c)
					cats = append(cats, c)
				} else {
					var p logic.Project
					json.Unmarshal(mapData, &p)
					proj = append(proj, p)
				}
			}
		case "yaml":
			for _, d := range data {
				mapData, err := yaml.Marshal(d)
				if err != nil {
					return err
				}
				if table == "categories" {
					var c logic.Category
					yaml.Unmarshal(mapData, &c)
					cats = append(cats, c)
				} else {
					var p logic.Project
					yaml.Unmarshal(mapData, &p)
					proj = append(proj, p)
				}
			}
		default:
			return errors.New("Incorrect format")
		}

		if table == "categories" {
			a = append(a, map[string]any{table: cats})
		} else {
			a = append(a, map[string]any{table: proj})
		}
	}
	if fileFormat == "yaml" {
		fileFormat = "yml"
	}
	file, fileErr := os.Create(fmt.Sprintf("%s/%s.%s", path, fileName, fileFormat))
	if fileErr != nil {
		return fileErr
	}
	defer file.Close()

	switch fileFormat {
	case "json":
		jsonData, err := json.MarshalIndent(a, "", "	")
		if err != nil {
			return err
		}
		file.Write(jsonData)
	case "yml":
		yamlData, err := yaml.Marshal(a)
		if err != nil {
			return err
		}
		file.Write(yamlData)
	}

	return nil
}
