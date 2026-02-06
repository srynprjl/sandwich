package initialize

import (
	"fmt"
	"strings"

	"github.com/srynprjl/stack/internal/category"
	"github.com/srynprjl/stack/internal/config"
	"github.com/srynprjl/stack/internal/projects"
)

func Init(lang string, p projects.Project) {
	exists, _ := p.Exists()
	if !exists {
		fmt.Println("Project not found! creating a new Project.")
		var name, shorthand, cats string
		var mapData = make(map[string]any)
		//taking input for new project
		fmt.Print("Name: \n>>")
		fmt.Scan(&name)
		fmt.Print("Shorthand: \n>>")
		if p.UID == "" {
			fmt.Scan(&shorthand)
		} else {
			fmt.Print(p.UID + "\n")
			shorthand = p.UID
		}
		fmt.Print("Category: \n>>")
		fmt.Scan(&cats)
		//selection later
		var c category.Category
		c = category.Category{UID: cats}
		cat := c.ID
		data, resp := c.GetField([]string{"id"})

		if resp.Error != nil || len(data) == 0 {
			fmt.Println(resp.Message)
			return
		}
		// fmt.Print(data["id"])
		cat = int(data["id"].(int64)) // error here if category doesnt exist
		mapData["name"] = name
		mapData["shorthand"] = shorthand
		mapData["category"] = cat
		p.Category = cat
		res := p.Add(mapData)
		if res.Status == 201 {
			fmt.Println("Created")
		}
	}
	// get project data
	data, res := p.Get()
	if res.Status != 200 {
		fmt.Println(res.Message)
		return
	}
	// update path with checks
	path := data["path"].(string)
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
		p.Update(map[string]any{"path": path})
	}

	if path == config.Conf.ProjectLocation || path == "/" {
		path = config.Conf.ProjectLocation + data["name"].(string) + "/"
		// update db tables
		p.Update(map[string]any{"path": path})
	}
	data["path"] = path
	// call functions based on lang
	switch lang {
	case "go", "golang":
		InitGo(data)
	case "js", "javascript":
	// code here
	case "java":
	// code here
	case "kotlin", "kt":
	// code here
	case "python", "py":
		InitPython(data)
	}
}
