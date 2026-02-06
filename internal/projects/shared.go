package projects

import (
	"errors"

	"github.com/srynprjl/stack/internal/utils/responses"
)

func (p *Project) checkIfCategoryExists() responses.Response {
	if p.Category == 0 {
		data, response := p.GetField([]string{"category"})
		if response.Error != nil {
			return response
		}
		if len(data) == 0 {
			return responses.Response{
				Message: "no project found associated with that id.",
				Error:   errors.New("no project found associated with that id."),
				Status:  200,
			}
		}
		p.Category = int(data["category"].(int64))
	}

	if exists, err := p.Exists(); !exists {
		if err != nil {
			return responses.Response{
				Message: err.Error(),
				Status:  500,
				Error:   err,
			}
		}
		return responses.Response{
			Message: "no category found associated with that id.",
			Status:  400,
			Error:   errors.New("no category found associated with that id."),
		}
	}

	return responses.Response{
		Error:  nil,
		Status: 200,
	}

}

func (p Project) makeConditions() map[string]any {
	conditions := make(map[string]any)
	if p.ProjectId != "" {
		conditions["shorthand"] = p.ProjectId
	}
	if p.Uuid != "" {
		conditions["uuid"] = p.Uuid
	}

	return conditions
}
