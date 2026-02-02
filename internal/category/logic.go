package category

import (
	"github.com/srynprjl/sandwich/utils/db"
)

func (c *Category) DoesExists() (bool, error) {
	return db.DB.CheckExists("categories", map[string]any{"id": c.Id})
}

func (c *Category) Add(data map[string]any) map[string]any {
	err := db.DB.InsertOne("categories", data)
	if err != nil {
		return map[string]any{"message": err.Error(), "status": "500"}
	}
	return map[string](any){"message": "Success: Inserted data", "status": "201", "data": *c}
}

func (c *Category) Delete() map[string]any {

	exists, existErr := c.DoesExists()
	if existErr != nil {
		return map[string]any{"message": existErr.Error(), "status": "400"}
	}
	if !exists {
		return map[string]any{"message": "The category doesn't exist", "status": "400"}
	}

	err := db.DB.DeleteItem("categories", map[string]any{"id": c.Id})
	if err != nil {
		return map[string]any{"message": err.Error(), "status": "500"}
	}
	return map[string]any{"message": "Deleted item", "status": "200"}
}

func (c *Category) Update(updateItems map[string]any) map[string]any {
	exists, existErr := c.DoesExists()
	if existErr != nil {
		return map[string]any{"message": existErr.Error(), "status": "400"}
	}
	if !exists {
		return map[string]any{"message": "The category doesn't exist", "status": "400"}
	}
	err := db.DB.UpdateItems("categories", updateItems, map[string]any{"id": c.Id})
	if err != nil {
		return map[string]any{"message": err.Error(), "status": "500"}
	}
	return map[string]any{"message": "Updated successfully", "status": "200"}
}

func (c *Category) GetField(field []string) map[string]any {
	data, err := db.DB.Query("categories", field, map[string]any{"id": c.Id})
	if err != nil {
		return map[string]any{"message": err.Error(), "status": 500, "data": map[string]any{}}
	}
	return map[string]any{"message": err.Error(), "status": 500, "data": data[0]}
}

func GetAll() map[string]any {
	conn := db.DB
	conn.Connect()
	defer conn.Conn.Close()
	data, err := db.DB.Query("categories", []string{}, map[string]any{})
	if err != nil {
		return map[string]any{"message": err.Error(), "status": "400", "data": []Category{}}
	}

	return map[string]any{"message": "Updated successfully", "status": "200", "data": data}
}
