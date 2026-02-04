package exports

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/srynprjl/sandwich/internal/utils/db"
	"gopkg.in/yaml.v3"
)

func Import(fileFormat string, file string) error {

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

	switch fileFormat {
	case "json":
		err := json.Unmarshal(data, &mapData)
		if err != nil {
			fmt.Println(err.Error())
		}
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
	case "yaml":
		err := yaml.Unmarshal(data, &mapData)
		if err != nil {
			fmt.Println(err.Error())
		}

		for _, data := range mapData {

			for k, v := range data {
				var value []map[string]any
				byteValues, err := yaml.Marshal(v)
				if err != nil {
					return err
				}
				yaml.Unmarshal(byteValues, &value)
				derr := db.DB.InsertMany(k, value)
				if derr != nil {
					fmt.Println(derr.Error())
				}
			}
		}
	default:
		return errors.New("Wrong file format")
	}

	return nil
}
