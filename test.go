package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
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
	data := projects.GetNRandom(1)["data"].([]projects.Project)
	fmt.Println(data[0].Id)
}
