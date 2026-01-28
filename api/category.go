package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/srynprjl/sandwich/internal/category"
	"github.com/srynprjl/sandwich/internal/projects"
)

func CategoryGetAll(r http.ResponseWriter, w *http.Request) {

	r.Header().Set("Content-Type", "application/json")
	resp := category.GetAll()
	if resp["status"] == "200" {
		data := resp["data"].([]category.Category)
		json.NewEncoder(r).Encode(data)
	} else {
		fmt.Fprintf(r, resp["message"].(string))
	}

}

func CategoryAdd(r http.ResponseWriter, w *http.Request) {

	var c category.Category
	json.NewDecoder(w.Body).Decode(&c)
	resp := c.Add()
	if resp["status"] == "201" {
		fmt.Fprintf(r, "Created")
	} else {
		fmt.Fprintf(r, "Failed")
	}

}

func CategoryDelete(r http.ResponseWriter, w *http.Request) {

	id, err := strconv.Atoi(w.PathValue("id"))
	if err != nil {
		fmt.Fprint(r, err.Error())
	}
	c := category.Category{Id: id}
	resp := c.Delete()
	if resp["status"] != "200" {
		fmt.Fprint(r, resp["message"])
		return
	}
	fmt.Fprint(r, "Deleted")

}

func CategoryUpdate(r http.ResponseWriter, w *http.Request) {
	id, err := strconv.Atoi(w.PathValue("id"))
	if err != nil {
		fmt.Printf("Error!")
		return
	}
	c := category.Category{Id: id}
	json.NewDecoder(w.Body).Decode(&c)
	resp := c.Update()
	fmt.Fprint(r, resp["message"])

}

func CategoryGetAllProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		fmt.Fprintf(w, "Error")
		return
	}
	c := category.Category{Id: id}
	resp := projects.GetProjects(c)
	if resp["status"] != "200" {
		fmt.Fprintln(w, resp["message"])
		return
	}
	json.NewEncoder(w).Encode(resp["data"].([]projects.Project))
}
