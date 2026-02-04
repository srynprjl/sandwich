package cmd

import (
	"strconv"

	"github.com/srynprjl/sandwich/internal/category"
	"github.com/srynprjl/sandwich/internal/projects"
)

func getCategoryForCondition(args []string) category.Category {
	if data, err := strconv.Atoi(args[0]); err == nil {
		return category.Category{Id: data}
	}
	return category.Category{Shorthand: args[0]}
}

func getProjectsForCondition(args []string) projects.Project {
	if data, err := strconv.Atoi(args[0]); err == nil {
		return projects.Project{Id: data}
	}
	return projects.Project{ProjectId: args[0]}
}
