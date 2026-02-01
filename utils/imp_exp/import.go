package impexp

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/srynprjl/sandwich/internal/category"
	"github.com/srynprjl/sandwich/internal/projects"
	"github.com/srynprjl/sandwich/utils/db"
)

func Import(fileFormat string, file string) error {
	// check fileFormat
	if fileFormat != "json" {
		return errors.New("Wrong file format")
	}
	// check if fileExists
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		log.Fatal(err)
	}

	// code to extract and run in database
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	var mapData []map[string]any
	json.Unmarshal(data, &mapData)
	for _, data := range mapData {
		for k, v := range data {
			var value []map[string]any
			byteValues, err := json.Marshal(v)
			if err != nil {
				return err
			}
			json.Unmarshal(byteValues, &value)
			derr := db.DB.InsertMany(k, value)
			if derr != nil {
				return derr
			}
		}
	}
	return errors.New("this should happen?")
}

func Export(fileFormat string, path string, fileName string, tables ...string) error {

	// check if file format correct\
	if fileFormat != "json" {
		return errors.New("Incorrect format")
	}
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
		var cats []category.Category
		var proj []projects.Project

		// take data from that table and map to the map or struct idk
		for _, d := range data {
			mapData, err := json.Marshal(d)
			if err != nil {
				return err
			}
			if table == "categories" {
				var c category.Category
				json.Unmarshal(mapData, &c)
				cats = append(cats, c)
			} else {
				var p projects.Project
				json.Unmarshal(mapData, &p)
				proj = append(proj, p)
			}
		}

		if table == "categories" {
			a = append(a, map[string]any{table: cats})
		} else {
			a = append(a, map[string]any{table: proj})
		}
	}

	file, fileErr := os.Create(fmt.Sprintf("%s/%s.%s", path, fileName, fileFormat))
	if fileErr != nil {
		return fileErr
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(a, "", "	")
	if err != nil {
		return err
	}
	// write data to a file
	file.Write(jsonData)
	// json.Unmarshal()
	return nil
}
