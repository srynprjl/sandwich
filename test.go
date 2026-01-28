package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/srynprjl/sandwich/internal/category"
	"github.com/srynprjl/sandwich/internal/projects"
)

func test() {
	// data := category.GetAll()["data"].([]category.Category)
	// for _, d := range data {
	// 	fmt.Println(d.Id, d.Title, d.Shorthand)
	// }
	// project := projects.Project{
	// 	Id:       5,
	// 	Category: 1,
	// }

	// res, _ := project.Exists()
	// fmt.Print(res)
	//
	cat := category.Category{Id: 2}
	projects.GetProjects(cat)

	// fmt.Println(project.Update(map[string]any{"name": "1234", "sdfasef": "3432", "category": 2}))
}
