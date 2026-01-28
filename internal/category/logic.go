package category

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/srynprjl/sandwich/utils"
)

func whereClause(id int, shorthand string) (string, []any) {
	var where_clause []string
	var values []any
	if id != 0 {
		where_clause = append(where_clause, "id = ?")
		values = append(values, id)
	}
	if shorthand != "" {
		where_clause = append(where_clause, "shorthand = ?")
		values = append(values, shorthand)
	}
	where := strings.Join(where_clause, " AND ")
	return where, values
}

func (c *Category) DoesExists() (bool, error) {
	conn := utils.DB
	conn.Connect()
	defer conn.Close()
	where, values := whereClause(c.Id, c.Shorthand)
	sqlS := "SELECT 1 FROM categories WHERE " + where + " LIMIT 1;"
	// fmt.Println(sqlS, values)
	var enough bool
	err := conn.Conn.QueryRow(sqlS, values...).Scan(&enough)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (c *Category) Add() map[string]any {
	conn := utils.DB
	conn.Connect()
	defer conn.Close()
	sql := `
		INSERT INTO categories(name, shorthand) VALUES(?, ?)
	`
	_, err := conn.Conn.Exec(sql, c.Title, c.Shorthand)
	if err != nil {
		return map[string]any{"message": "Failed to add data. Internal Server Error", "status": "500"}
	}
	return map[string](any){"message": "Successfully added data", "status": "201", "data": *c}

}

func (c *Category) Delete() map[string]any {
	conn := utils.DB
	conn.Connect()
	defer conn.Close()
	exists, existErr := c.DoesExists()
	if existErr != nil {
		return map[string]any{"message": existErr.Error(), "status": "400"}
	}
	if !exists {
		return map[string]any{"message": "The category doesn't exist", "status": "400"}
	}
	where, values := whereClause(c.Id, c.Shorthand)
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
	exists, existErr := c.DoesExists()
	if existErr != nil {
		return map[string]any{"message": existErr.Error(), "status": "400"}
	}
	if !exists {
		return map[string]any{"message": "The category doesn't exist", "status": "400"}
	}
	var update_values []string
	var values []any
	if c.Title != "" {
		update_values = append(update_values, "name = ?")
		values = append(values, c.Title)
	}
	if c.Shorthand != "" {
		update_values = append(update_values, "shorthand = ?")
		values = append(values, c.Shorthand)
	}
	if c.Id == 0 {
		return map[string]any{"message": "ID Needed", "status": "400"}
	}
	set := strings.Join(update_values, ",")
	sql := fmt.Sprintf(`UPDATE categories SET %s WHERE id=%d`, set, c.Id)

	_, err := conn.Conn.Exec(sql, values...)
	if err != nil {
		return map[string]any{"message": err.Error(), "status": "500"}
	}
	return map[string]any{"message": "Updated successfully", "status": "200"}
}

func (c *Category) GetProjects() {
	// pass
}

func GetAll() map[string]any {
	conn := utils.DB
	conn.Connect()
	defer conn.Conn.Close()
	sqlStatement := "SELECT * FROM categories"
	rows, err := conn.Conn.Query(sqlStatement)
	if err != nil {
		return map[string]any{"message": err.Error(), "status": "400", "data": []Category{}}
	}
	data := []Category{}
	for rows.Next() {
		c := &Category{}
		err := rows.Scan(&c.Id, &c.Title, &c.Shorthand)
		if err != nil {
			return map[string]any{"message": err.Error(), "status": "400", "data": []Category{}}
		}
		data = append(data, *c)
	}
	return map[string]any{"message": "Updated successfully", "status": "200", "data": data}
}
