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
	project := projects.Project{
		Id: 5,
	}

	// res, _ := project.Exists()
	// fmt.Print(res)
	//

	data := project.GetField([]string{"name", "description"})
	fmt.Println(data)
}
