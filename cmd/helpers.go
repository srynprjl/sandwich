package cmd

import (
	"strconv"

	"github.com/srynprjl/stack/internal/category"
	"github.com/srynprjl/stack/internal/projects"
)

func GetCategoryForCondition(args []string) category.Category {
	if data, err := strconv.Atoi(args[0]); err == nil {
		return category.Category{ID: data}
	}
	return category.Category{UID: args[0]}
}

func GetProjectsForCondition(args []string) projects.Project {
	if data, err := strconv.Atoi(args[0]); err == nil {
		return projects.Project{ID: data}
	}
	return projects.Project{UID: args[0]}
}
