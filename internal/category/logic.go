package category

import (
	"fmt"
	"strings"

	"github.com/srynprjl/sandwich/utils"
)

func (c *Category) Add() map[string]any {
	conn := utils.DB
	conn.Connect()
	defer conn.Close()
	sql := `
		INSERT INTO categories(name, shorthand) VALUES(?, ?)
	`
	_, err := conn.Conn.Exec(sql, *c.Title, *c.Shorthand)
	if err != nil {
		return map[string]any{"message": "Failed to add data. Internal Server Error", "status": "500"}
	}
	return map[string](any){"message": "Successfully added data", "status": "201", "data": *c}

}

func (c *Category) Delete() map[string]any {
	conn := utils.DB
	conn.Connect()
	defer conn.Close()
	var where_clause []string
	var values []any
	if c.Id != nil {
		where_clause = append(where_clause, "id = ?")
		values = append(values, *c.Id)
	}
	if c.Shorthand != nil {
		where_clause = append(where_clause, "shorthand = ?")
		values = append(values, *c.Shorthand)
	}
	where := strings.Join(where_clause, " AND ")
	sql := "DELETE FROM categories WHERE " + where + ";"
	if len(where) <= 0 {
		return map[string]any{"message": "No ID or Shorthand Provided", "status": "400"}
	}
	_, err := conn.Conn.Exec(sql, values...)
	if err != nil {
		return map[string]any{"message": err.Error(), "status": "500"}
	}
	return map[string]any{"message": "Deleted item", "status": "200"}
}

func (c *Category) Update() map[string]any {
	conn := utils.DB
	conn.Connect()
	defer conn.Conn.Close()
	var update_values []string
	var values []any
	if c.Title != nil {
		update_values = append(update_values, "name = ?")
		values = append(values, *c.Title)
	}
	if c.Shorthand != nil {
		update_values = append(update_values, "shorthand = ?")
		values = append(values, *c.Shorthand)
	}
	if c.Id == nil {
		return map[string]any{"message": "ID Needed", "status": "400"}
	}
	set := strings.Join(update_values, ",")
	sql := fmt.Sprintf(`UPDATE categories SET %s WHERE id=%d`, set, *c.Id)

	_, err := conn.Conn.Exec(sql, values...)
	if err != nil {
		return map[string]any{"message": err.Error(), "status": "500"}
	}
	return map[string]any{"message": "Updated successfully", "status": "200"}
}

func (c *Category) GetProject() {

}

func GetAll() {

}
