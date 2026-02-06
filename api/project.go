package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/srynprjl/stack/internal/category"
	"github.com/srynprjl/stack/internal/projects"
)

func ProjectGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")
	catUUID := r.PathValue("catId")
	c := category.Category{UUID: catUUID}
	d, res := c.GetField([]string{"id"})
	if res.Error != nil {
		json.NewEncoder(w).Encode(res.WebResponse())
	}
	catId := int(d["id"].(int64))

	project := projects.Project{UUID: id, Category: catId}
	data, resp := project.Get()
	response := resp.WebResponse()
	response["data"] = data
	json.NewEncoder(w).Encode(response)
}

func ProjectGetRandom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, resp := projects.GetRandom(1)
	response := resp.WebResponse()
	response["data"] = data
	json.NewEncoder(w).Encode(response)
}

func ProjectGetNRandom(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.Atoi(r.PathValue("num"))
	if err != nil {
		fmt.Fprintln(w, "ERROR")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	data, resp := projects.GetRandom(n)
	response := resp.WebResponse()
	response["data"] = data
	json.NewEncoder(w).Encode(response)
}

func ProjectAdd(w http.ResponseWriter, r *http.Request) {
	catUUID := r.PathValue("catId")
	c := category.Category{UUID: catUUID}
	d, res := c.GetField([]string{"id"})
	if res.Error != nil {
		json.NewEncoder(w).Encode(res.WebResponse())
	}
	catId := int(d["id"].(int64))

	p := projects.Project{Category: catId}
	data := make(map[string]any)
	json.NewDecoder(r.Body).Decode(&data)
	data["category"] = p.Category
	resp := p.Add(data)
	json.NewEncoder(w).Encode(resp.WebResponse())
}

func ProjectDelete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	catUUID := r.PathValue("catId")

	c := category.Category{UUID: catUUID}
	d, res := c.GetField([]string{"id"})
	if res.Error != nil {
		json.NewEncoder(w).Encode(res.WebResponse())
	}
	catId := int(d["id"].(int64))

	p := projects.Project{UUID: id, Category: catId}
	resp := p.Remove()
	json.NewEncoder(w).Encode(resp.WebResponse())
}

func ProjectUpdate(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	catUUID := r.PathValue("catId")

	c := category.Category{UUID: catUUID}
	d, res := c.GetField([]string{"id"})
	if res.Error != nil {
		json.NewEncoder(w).Encode(res.WebResponse())
	}
	catId := int(d["id"].(int64))

	p := projects.Project{UUID: id, Category: catId}

	var data map[string]any

	jsonErr := json.NewDecoder(r.Body).Decode(&data)
	if jsonErr != nil {
		fmt.Fprintln(w, "Error!")
	}

	resp := p.Update(data)
	json.NewEncoder(w).Encode(resp.WebResponse())
}
