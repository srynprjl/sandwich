package category

import (
	"errors"

	"github.com/srynprjl/stack/internal/utils/db"
	"github.com/srynprjl/stack/internal/utils/responses"
)

func (c *Category) getConditions() map[string]any {
	conditions := make(map[string]any)
	if c.ID != 0 {
		conditions["id"] = c.ID
	}
	if c.UID != "" {
		conditions["shorthand"] = c.UID
	}
	if c.UUID != "" {
		conditions["uuid"] = c.UUID
	}

	return conditions
}

func (c *Category) DoesExists() (bool, error) {
	conditions := c.getConditions()
	return db.DB.CheckExists("categories", conditions)
}

func (c *Category) Add(data map[string]any) responses.Response {
	err := db.DB.InsertOne("categories", data)
	if err != nil {
		return responses.Response{
			Status:  500,
			Message: err.Error(),
			Error:   err,
		}
	}
	return responses.Response{
		Status:  201,
		Message: "Successfully added data!",
		Error:   nil,
	}
}

func (c *Category) Delete() responses.Response {
	exists, existErr := c.DoesExists()
	if existErr != nil {
		return responses.Response{
			Status:  400,
			Message: existErr.Error(),
			Error:   existErr,
		}
	}
	if !exists {
		return responses.Response{
			Status:  400,
			Message: "category doesn't exist",
			Error:   errors.New("category doesn't exist"),
		}
	}
	conditions := c.getConditions()
	err := db.DB.DeleteItem("categories", conditions)
	if err != nil {
		return responses.Response{
			Status:  500,
			Message: err.Error(),
			Error:   err,
		}
	}
	return responses.Response{
		Status:  200,
		Message: "successfully deleted item.",
		Error:   nil,
	}
}

func (c *Category) Update(updateItems map[string]any) responses.Response {
	exists, existErr := c.DoesExists()
	if existErr != nil {
		return responses.Response{
			Status:  400,
			Message: existErr.Error(),
			Error:   existErr,
		}
	}
	if !exists {
		return responses.Response{
			Status:  400,
			Message: "the category doesnt exist",
			Error:   errors.New("the category doesnt exist"),
		}
	}
	conditions := c.getConditions()
	err := db.DB.UpdateItems("categories", updateItems, conditions)
	if err != nil {
		return responses.Response{
			Status:  500,
			Message: err.Error(),
			Error:   err,
		}
	}
	return responses.Response{
		Status:  200,
		Message: "category update successfully",
		Error:   existErr,
	}
}

func (c *Category) GetField(field []string) (map[string]any, responses.Response) {
	conditions := c.getConditions()
	data, err := db.DB.Query("categories", field, conditions)
	if err != nil {
		return map[string]any{}, responses.Response{Status: 500, Message: err.Error(), Error: err}
	}
	if len(data) == 0 {
		return map[string]any{}, responses.Response{Status: 200, Message: "No field found.", Error: nil}
	}
	return data[0], responses.Response{
		Status:  200,
		Message: "fetched data successfully",
		Error:   nil,
	}
}

func GetAll() ([]map[string]any, responses.Response) {
	data, err := db.DB.Query("categories", []string{}, map[string]any{})
	if err != nil {
		return []map[string]any{}, responses.Response{
			Status:  500,
			Message: err.Error(),
			Error:   err,
		}
	}

	return data, responses.Response{
		Status:  200,
		Message: "data fetched succesfully",
		Error:   nil,
	}
}
