package category

import (
	"github.com/srynprjl/sandwich/internal/utils/db"
)

func (c *Category) getConditions() map[string]any {
	conditions := make(map[string]any)
	if c.Id != 0 {
		conditions["id"] = c.Id
	}
	if c.Shorthand != "" {
		conditions["shorthand"] = c.Shorthand
	}
	if c.Uuid != "" {
		conditions["uuid"] = c.Uuid
	}

	return conditions
}

func (c *Category) DoesExists() (bool, error) {
	conditions := c.getConditions()
	return db.DB.CheckExists("categories", conditions)
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
	conditions := c.getConditions()
	err := db.DB.DeleteItem("categories", conditions)
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
	conditions := c.getConditions()
	err := db.DB.UpdateItems("categories", updateItems, conditions)
	if err != nil {
		return map[string]any{"message": err.Error(), "status": "500"}
	}
	return map[string]any{"message": "Updated successfully", "status": "200"}
}

func (c *Category) GetField(field []string) map[string]any {
	conditions := c.getConditions()
	data, err := db.DB.Query("categories", field, conditions)
	if err != nil {
		return map[string]any{"message": err.Error(), "status": 500, "data": map[string]any{}}
	}
	return map[string]any{"message": "Succcess", "status": 200, "data": data[0]}
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
