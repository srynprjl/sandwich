package projects

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/srynprjl/stack/internal/category"
	"github.com/srynprjl/stack/internal/config"
	"github.com/srynprjl/stack/internal/utils/db"
	"github.com/srynprjl/stack/internal/utils/responses"
)

func (p *Project) Exists() (bool, error) {
	if p.ProjectId == "" {
		d, res := p.GetField([]string{"shorthand"})
		if res.Error != nil {
			return false, res.Error
		}
		if len(d) == 0 {
			return false, errors.New("project doesn't exist")
		}
		p.ProjectId = d["shorthand"].(string)
	}
	if p.Category == 0 {
		d, res := p.GetField([]string{"category"})
		if res.Error != nil {
			return false, res.Error
		}
		if len(d) == 0 {
			return false, errors.New("project doesn't exist")
		}
		p.Category = int(d["category"].(int64))
	}
	exists, err := db.DB.CheckExists("projects", map[string]any{"shorthand": p.ProjectId, "category": p.Category})
	if err != nil {
		if err == sql.ErrNoRows {
			return exists, nil
		}
		return false, err
	}

	return exists, nil
}

func (p *Project) Add(insertData map[string]any) responses.Response {
	c := category.Category{Id: p.Category}
	if exists, err := c.DoesExists(); !exists {
		if err != nil {
			return responses.Response{
				Error:   err,
				Message: err.Error(),
				Status:  500,
			}
		}
		return responses.Response{
			Error:   errors.New("category doesn't exist"),
			Message: "category doesn't exist",
			Status:  400,
		}
	}

	err := db.DB.InsertOne("projects", insertData)
	if err != nil {
		return responses.Response{
			Error:   err,
			Message: err.Error(),
			Status:  500,
		}
	}
	return responses.Response{
		Error:   nil,
		Message: "project successfully created.",
		Status:  201,
	}
}

func (p *Project) Remove() responses.Response {
	resp := p.checkIfCategoryExists()
	if resp.Error != nil {
		return resp
	}

	conditions := p.makeConditions()
	conditions["category"] = p.Category
	deleteConditions := conditions

	err := db.DB.DeleteItem("projects", deleteConditions)
	if err != nil {
		return responses.Response{
			Message: err.Error(),
			Error:   err,
			Status:  500,
		}
	}
	return responses.Response{
		Message: "project deleted succesfully",
		Error:   nil,
		Status:  200,
	}
}

func schemaUpdate(values map[string]any) map[string]any {
	keys := config.DefaultTables["projects"].Columns[2:]
	newData := make(map[string]any)
	for _, data := range keys {
		if value, exists := values[data]; exists {
			newData[data] = value
		}
	}
	return newData
}

func (p *Project) Update(updateData map[string]any) responses.Response {
	resp := p.checkIfCategoryExists()
	if resp.Error != nil {
		return resp
	}

	validatedData := schemaUpdate(updateData)
	if len(validatedData) == 0 {
		return responses.Response{
			Message: "No field to update",
			Status:  400,
			Error:   nil,
		}
	}
	err := db.DB.UpdateItems("projects", validatedData, map[string]any{"id": p.Id})
	if err != nil {
		return responses.Response{
			Error:   err,
			Message: err.Error(),
			Status:  500,
		}
	}
	return responses.Response{
		Error:   nil,
		Message: "data updated successfully",
		Status:  200,
	}
}

func (p *Project) Get() (map[string]any, responses.Response) {
	resp := p.checkIfCategoryExists()
	if resp.Error != nil {
		return map[string]any{}, resp
	}
	data, err := db.DB.QueryLimit("projects", []string{}, map[string]any{"shorthand": p.ProjectId, "category": p.Category}, 1)
	if err != nil {
		return map[string]any{}, responses.Response{
			Message: err.Error(),
			Status:  500,
			Error:   err,
		}
	}
	return data[0], responses.Response{
		Message: "projects fetched successfully",
		Status:  200,
		Error:   nil,
	}
}

func GetRandom(n int) ([]map[string]any, responses.Response) {
	data, err := db.DB.QueryRandom("projects", []string{}, map[string]any{}, n)
	if err != nil {
		return []map[string]any{}, responses.Response{
			Message: err.Error(),
			Status:  500,
			Error:   err,
		}
	}
	return data, responses.Response{
		Message: "projects fetched successfully",
		Status:  200,
		Error:   nil,
	}
}

func (p *Project) GetField(field []string) (map[string]any, responses.Response) {
	conditions := p.makeConditions()
	data, err := db.DB.Query("projects", field, conditions)
	if len(data) <= 0 {
		return map[string]any{}, responses.Response{
			Message: "No data found",
			Status:  200,
			Error:   nil,
		}
	}
	if err != nil {
		return map[string]any{}, responses.Response{
			Message: err.Error(),
			Status:  500,
			Error:   err,
		}
	}
	return data[0], responses.Response{
		Message: "project fetched succesfully",
		Status:  200,
		Error:   nil,
	}
}

func GetProjects(c category.Category) ([]map[string]any, responses.Response) {
	if exists, err := c.DoesExists(); !exists {
		if err != nil {
			return []map[string]any{}, responses.Response{
				Message: err.Error(),
				Status:  500,
				Error:   err,
			}
		}
		return []map[string]any{}, responses.Response{
			Message: "No category found",
			Status:  500,
			Error:   errors.New("category not found"),
		}
	}
	if c.Id == 0 {
		data, resp := c.GetField([]string{"id"})
		if resp.Error != nil {
			fmt.Println(resp.Message)
			return []map[string]any{}, resp
		}
		c.Id = int(data["id"].(int64))
	}

	data, err := db.DB.Query("projects", []string{}, map[string]any{"category": c.Id})
	if err != nil {
		return []map[string]any{}, responses.Response{
			Message: err.Error(),
			Status:  500,
			Error:   err,
		}
	}
	return data, responses.Response{Status: 200, Message: "projects fetched successfully", Error: nil}
}

func GetProjectWhere(condition map[string]any) ([]map[string]any, responses.Response) {
	data, err := db.DB.Query("projects", []string{}, condition)
	if err != nil {
		return []map[string]any{}, responses.Response{Status: 500, Message: err.Error(), Error: err}
	}
	return data, responses.Response{Status: 200, Message: "projects fetched successfully", Error: nil}
}
