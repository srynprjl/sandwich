package init

import (
	"fmt"
	"os"
	"strconv"

	"github.com/srynprjl/sandwich/internal/category"
	"github.com/srynprjl/sandwich/internal/config"
	"github.com/srynprjl/sandwich/internal/projects"
)

func Init(lang string, p *projects.Project) {
	fmt.Println("Support for " + lang + " is avalaible")
	fmt.Println(p)
	exists, _ := p.Exists()

	// if project doesnt exist create a project with that id
	if !exists {
		fmt.Println("Project not found! creating a new Project.")
		var name, shorthand, cats string
		var mapData = make(map[string]any)
		//taking input for new project
		fmt.Print("Name: ")
		fmt.Scan(&name)
		fmt.Print("Shorthand: ")
		if p.ProjectId == "" {
			fmt.Scan(&shorthand)
		} else {
			fmt.Print(p.ProjectId + "\n")
			shorthand = p.ProjectId
		}
		fmt.Print("Category: ")
		fmt.Scan(&cats)
		var c category.Category
		if data, err := strconv.Atoi(cats); err == nil {
			c = category.Category{Id: data}
		} else {
			c = category.Category{Shorthand: cats}
		}
		var data map[string]any
		cat := c.Id
		if cat == 0 {
			data = c.GetField([]string{"id"})
			if data["status"] != "200" {
				fmt.Println(data["message"])
				return
			}
			cat = int(data["data"].(map[string]any)["id"].(int64))
		}
		mapData["name"] = name
		mapData["shorthand"] = shorthand
		mapData["category"] = cat
		p.Category = cat
		res := p.Add(mapData)
		if res["status"] == "201" {
			fmt.Println("Created")
		}
	}
	// create folder with checks
	res := p.Get()
	if res["status"] != "200" {
		fmt.Println(res["message"])
		return
	}
	project_data := res["data"].(map[string]any)
	if project_data["path"] == config.Conf.ProjectLocation || project_data["path"] == "" {
		os.MkdirAll(config.Conf.ProjectLocation+project_data["name"].(string), 0700)
		project_data["path"] = config.Conf.ProjectLocation + project_data["name"].(string)
		// update db tables
		p.Update(map[string]any{"path": project_data["path"]})
	}
	
	// call functions based on lang
	switch lang {
	case "test":
		// code here
	}
}
