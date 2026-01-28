package projects

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/srynprjl/sandwich/internal/category"
	"github.com/srynprjl/sandwich/utils"
)

func (p *Project) MapProject() map[string]any {
	c := category.Category{Id: p.Category}
	data := c.GetField([]string{"name"})["data"].(string)
	return map[string]any{"id": p.Id, "name": p.Title, "description": p.Description, "path": p.Path, "category": data, "favourite": p.Favourite, "completed": p.Completed}
}

func (p *Project) Exists() (bool, error) {
	var exists bool = false
	db := utils.DB
	db.Connect()
	defer db.Close()

	if p.Id <= 0 || p.Category <= 0 {
		return false, errors.New("Id/Category should be  given.")
	}

	query := "SELECT 1 FROM projects WHERE id = ? AND category = ? LIMIT 1"
	err := db.Conn.QueryRow(query, p.Id, p.Category).Scan(&exists)

	if err != nil {
		if err == sql.ErrNoRows {
			return exists, nil
		}
		return false, err
	}

	return exists, nil
}

func (p *Project) Add() map[string]any {
	db := utils.DB
	db.Connect()
	defer db.Close()
	sqlStatement := "INSERT INTO projects(name, description, completed, favorite, path, category) VALUES (?, ?, ?, ?, ?, ?)"
	values := []any{p.Title, p.Description, p.Completed, p.Favourite, p.Path, p.Category}
	_, err := db.Conn.Exec(sqlStatement, values...)
	if err != nil {
		return map[string]any{"message": "Failed.", "status": "500"}
	}
	return map[string]any{"message": "Created.", "status": "201"}
}

func (p *Project) Remove() map[string]any {
	db := utils.DB
	db.Connect()
	defer db.Close()
	sqlStatement := "DELETE FROM projects WHERE id=? AND category=?"
	values := []any{p.Id, p.Category}
	if exists, err := p.Exists(); !exists {
		if err != nil {
			return map[string]any{"message": "Failed.", "status": "500"}
		}
		return map[string]any{"message": "No project found in that category", "status": "400"}
	}
	_, err := db.Conn.Exec(sqlStatement, values...)
	if err != nil {
		return map[string]any{"message": "Failed.", "status": "500"}
	}
	return map[string]any{"message": "Deleted.", "status": "201"}
}

func schemaUpdate(values map[string]any) map[string]any {
	keys := []string{"name", "description", "completed", "favorite", "path", "category"}
	newData := make(map[string]any)
	for _, data := range keys {
		if value, exists := values[data]; exists {
			newData[data] = value
		}
	}
	return newData
}

func (p *Project) Update(newValues map[string]any) map[string]any {
	db := utils.DB
	db.Connect()
	defer db.Close()
	if exists, err := p.Exists(); !exists {
		if err != nil {
			return map[string]any{"message": "Failed.", "status": "500"}
		}
		return map[string]any{"message": "No project found in that category", "status": "400"}
	}
	validatedData := schemaUpdate(newValues)
	if len(validatedData) == 0 {
		return map[string]any{"message": "Nothing to be updated", "status": "200"}
	}
	var updateString []string
	var appendItems []any
	for k, v := range validatedData {
		updateString = append(updateString, fmt.Sprintf("%s=?", k))
		appendItems = append(appendItems, v)
	}
	appendItems = append(appendItems, p.Id)
	updateField := strings.Join(updateString, ",")
	sql := fmt.Sprintf("UPDATE projects SET %s WHERE id = ?; ", updateField)
	_, err := db.Conn.Exec(sql, appendItems...)
	if err != nil {
		return map[string]any{"message": err.Error(), "status": "500"}
	}
	return map[string]any{"message": "Updated", "status": "200"}

}

func (p *Project) Get() map[string]any {
	var fields ProjectFields
	fields.Init(p)
	db := utils.DB
	db.Connect()
	defer db.Close()
	id := p.Id
	cat := p.Category
	if exists, err := p.Exists(); !exists {
		if err != nil {
			return map[string]any{"message": "Failed.", "status": "500"}
		}
		return map[string]any{"message": "No project found in that category", "status": "400"}
	}
	sqlStatement := "SELECT * FROM projects WHERE id=? AND category=? LIMIT 1"
	err := db.Conn.QueryRow(sqlStatement, id, cat).Scan(fields.Field...)
	if err != nil {
		return map[string]any{"message": err.Error(), "status": "500"}
	}
	return map[string]any{"message": "Fetched.", "data": *p, "status": "200"}
}

func GetRandom() map[string]any {
	var p Project
	db := utils.DB
	db.Connect()
	defer db.Close()
	var fields ProjectFields
	fields.Init(&p)
	sqlStatement := "SELECT * FROM projects ORDER BY RANDOM() LIMIT 1"
	err := db.Conn.QueryRow(sqlStatement).Scan(fields.Field...)
	if err != nil {
		return map[string]any{"message": err.Error(), "status": "500"}
	}
	return map[string]any{"message": "Fetched.", "data": p, "status": "200"}
}

func GetNRandom(n int) map[string]any {
	db := utils.DB
	db.Connect()
	defer db.Close()
	var projects []Project

	sqlStatement := fmt.Sprintf("SELECT * FROM projects ORDER BY RANDOM() LIMIT %d", n)
	rows, err := db.Conn.Query(sqlStatement)
	if err != nil {
		return map[string]any{"message": err.Error(), "status": "500"}
	}
	for rows.Next() {
		var p Project
		var fields ProjectFields
		fields.Init(&p)
		rows.Scan(fields.Field...)
		projects = append(projects, p)
	}
	return map[string]any{"message": "Fetched.", "data": projects, "status": "200"}
}

func (p *Project) GetField(field []string) map[string]any {
	db := utils.DB
	db.Connect()
	defer db.Close()
	id := p.Id
	values := strings.Join(field, ", ")
	query := fmt.Sprintf("SELECT %s FROM projects WHERE id= ?", values)
	var value = make([]any, len(field))
	var scanArgs = make([]any, len(field))
	for i := range value {
		scanArgs[i] = &value[i]
	}
	err := db.Conn.QueryRow(query, id).Scan(scanArgs...)
	if err != nil {
		return map[string]any{"message": err.Error(), "status": "500"}
	}
	return map[string]any{"message": "Fetched.", "data": value, "status": "200"}
}

func GetProjects(c category.Category) map[string]any {
	db := utils.DB
	db.Connect()
	defer db.Close()
	id := c.Id
	if exists, err := c.DoesExists(); !exists {
		if err != nil {
			return map[string]any{"message": err.Error(), "status": "500"}
		}
		return map[string]any{"message": "No categories found", "status": "500"}
	}
	query := "SELECT * FROM projects WHERE category = ?"
	res, err := db.Conn.Query(query, id)
	if err != nil {
		return map[string]any{"message": err.Error(), "status": "500"}
	}
	var project []Project
	for res.Next() {
		var p Project
		var pf ProjectFields
		pf.Init(&p)
		res.Scan(pf.Field...)
		project = append(project, p)
	}

	return map[string]any{"message": "Fetched", "data": project, "status": "200"}
}
