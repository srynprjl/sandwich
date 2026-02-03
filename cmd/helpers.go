package cmd

import (
	"strconv"

	"github.com/srynprjl/sandwich/internal/logic"
)

func getCategoryForCondition(args []string) (c logic.Category) {
	if data, err := strconv.Atoi(args[0]); err == nil {
		return logic.Category{Id: data}
	}
	return logic.Category{Shorthand: args[0]}
}

func getProjectsForCondition(args []string) (p logic.Project) {
	if data, err := strconv.Atoi(args[0]); err == nil {
		return logic.Project{Id: data}
	}
	return logic.Project{ProjectId: args[0]}
}
