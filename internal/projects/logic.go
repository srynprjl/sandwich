package projects

import (
	"database/sql"
	"errors"

	"github.com/srynprjl/sandwich/internal/category"
	"github.com/srynprjl/sandwich/utils/config"
	"github.com/srynprjl/sandwich/utils/db"
)

func (p *Project) Exists() (bool, error) {
	if p.Id <= 0 {
		return false, errors.New("Id should be  given.")
	}
	if p.Category == 0 {
		d := p.GetField([]string{"category"})["data"].(map[string]any)
		p.Category = int(d["category"].(int64))
	}
	exists, err := db.DB.CheckExists("projects", map[string]any{"id": p.Id, "category": p.Category})
	if err != nil {
		if err == sql.ErrNoRows {
			return exists, nil
		}
		return false, err
	}

	return exists, nil
}

func (p *Project) Add(insertData map[string]any) map[string]any {
	c := category.Category{Id: p.Category}
	if exists, err := c.DoesExists(); !exists {
		if err != nil {
			return map[string]any{"message": err.Error(), "status": "500"}
		}
		return map[string]any{"message": "Category doesn't exists", "status": "400"}
	}
	err := db.DB.InsertOne("projects", insertData)
	if err != nil {
		return map[string]any{"message": err.Error(), "status": "500"}
	}
	return map[string]any{"message": "Created.", "status": "201"}
}

func (p *Project) Remove() map[string]any {
	if p.Category == 0 {
		d := p.GetField([]string{"category"})["data"].(map[string]any)
		p.Category = int(d["category"].(int64))
	}
	if exists, err := p.Exists(); !exists {
		if err != nil {
			return map[string]any{"message": err.Error(), "status": "500"}
		}
		return map[string]any{"message": "No project found in that category", "status": "400"}
	}
	deleteConditions := map[string]any{"id": p.Id, "category": p.Category}

	err := db.DB.DeleteItem("projects", deleteConditions)
	if err != nil {
		return map[string]any{"message": err.Error(), "status": "500"}
	}
	return map[string]any{"message": "Deleted.", "status": "201"}
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

func (p *Project) Update(updateData map[string]any) map[string]any {
	if p.Category == 0 {
		d := p.GetField([]string{"category"})["data"].(map[string]any)
		p.Category = int(d["category"].(int64))
	}
	if exists, err := p.Exists(); !exists {
		if err != nil {
			return map[string]any{"message": "Failed.", "status": "500"}
		}
		return map[string]any{"message": "No project found in that category", "status": "400"}
	}
	validatedData := schemaUpdate(updateData)
	if len(validatedData) == 0 {
		return map[string]any{"message": "Nothing to be updated", "status": "200"}
	}
	err := db.DB.UpdateItems("projects", validatedData, map[string]any{"id": p.Id})
	if err != nil {
		return map[string]any{"message": err.Error(), "status": "500"}
	}
	return map[string]any{"message": "Updated data", "status": "200"}
}

func (p *Project) Get() map[string]any {
	if p.Category == 0 {
		d := p.GetField([]string{"category"})["data"].(map[string]any)
		p.Category = int(d["category"].(int64))
	}
	if exists, err := p.Exists(); !exists {
		if err != nil {
			return map[string]any{"message": "Failed.", "status": "500"}
		}
		return map[string]any{"message": "No project found in that category", "status": "400"}
	}
	data, err := db.DB.QueryLimit("projects", []string{}, map[string]any{"id": p.Id, "category": p.Category}, 1)
	if err != nil {
		return map[string]any{"message": err.Error(), "status": "500"}
	}
	return map[string]any{"message": "Fetched.", "data": data[0], "status": "200"}
}

func GetRandom() map[string]any {
	data, err := db.DB.QueryRandom("projects", []string{}, map[string]any{}, 1)
	if err != nil {
		return map[string]any{"message": err.Error(), "status": "500"}
	}
	return map[string]any{"message": "Fetched.", "data": data[0], "status": "200"}
}

func GetNRandom(n int) map[string]any {
	data, err := db.DB.QueryRandom("projects", []string{}, map[string]any{}, n)
	if err != nil {
		return map[string]any{"message": err.Error(), "status": "500"}
	}
	return map[string]any{"message": "Fetched.", "data": data, "status": "200"}
}

func (p *Project) GetField(field []string) map[string]any {
	data, err := db.DB.Query("projects", field, map[string]any{"id": p.Id})
	if err != nil {
		return map[string]any{"message": err.Error(), "status": "500", "data": map[string]any{}}
	}
	return map[string]any{"message": "Fetched.", "data": data[0], "status": "200"}
}

func GetProjects(c category.Category) map[string]any {
	if exists, err := c.DoesExists(); !exists {
		if err != nil {
			return map[string]any{"message": err.Error(), "status": "500"}
		}
		return map[string]any{"message": "No categories found", "status": "500"}
	}
	data, err := db.DB.Query("projects", []string{}, map[string]any{"category": c.Id})
	if err != nil {
		return map[string]any{"message": err.Error(), "status": "500"}
	}

	return map[string]any{"message": "Fetched", "data": data, "status": "200"}
}

func GetProjectWhere(condition map[string]any) map[string]any {
	data, err := db.DB.Query("projects", []string{}, condition)
	if err != nil {
		return map[string]any{"message": err.Error(), "status": "500"}
	}
	return map[string]any{"message": "Fetched", "data": data, "status": "200"}
}
