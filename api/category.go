package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/srynprjl/stack/internal/category"
	"github.com/srynprjl/stack/internal/projects"
)

func CategoryGetAll(r http.ResponseWriter, w *http.Request) {
	r.Header().Set("Content-Type", "application/json")
	data, resp := category.GetAll()
	response := resp.WebResponse()
	response["data"] = data
	json.NewEncoder(r).Encode(response)
}

func CategoryAdd(r http.ResponseWriter, w *http.Request) {
	var c category.Category
	var data = make(map[string]any)
	json.NewDecoder(w.Body).Decode(&data)
	resp := c.Add(data)
	json.NewEncoder(r).Encode(resp.WebResponse())
}

func CategoryDelete(r http.ResponseWriter, w *http.Request) {
	c := category.Category{UUID: w.PathValue("id")}
	resp := c.Delete()
	json.NewEncoder(r).Encode(resp.WebResponse())
}

func CategoryUpdate(r http.ResponseWriter, w *http.Request) {
	id := w.PathValue("id")
	c := category.Category{UUID: id}
	updateData := make(map[string]any)
	json.NewDecoder(w.Body).Decode(&updateData)
	resp := c.Update(updateData)
	json.NewEncoder(r).Encode(resp.WebResponse())

}

func CategoryGetAllProjects(r http.ResponseWriter, w *http.Request) {
	r.Header().Set("Content-Type", "application/json")
	id := w.PathValue("id")
	c := category.Category{UUID: id}
	data, res := projects.GetProjects(c)
	if res.Status != 200 {
		fmt.Printf("Error: %s\n", res.Message)
		return
	}
	response := res.WebResponse()
	response["data"] = data
	json.NewEncoder(r).Encode(response)
}
