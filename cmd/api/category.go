package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/srynprjl/sandwich/internal/category"
)

func CategoryGetAll(r http.ResponseWriter, w *http.Request) {
	if w.Method == "GET" {
		resp := category.GetAll()
		if resp["status"] == "200" {
			data := resp["data"].([]category.Category)
			json.NewEncoder(r).Encode(data)
		} else {
			fmt.Fprintf(r, resp["message"].(string))
		}
	}
}

func CategoryAdd(r http.ResponseWriter, w *http.Request) {
	if w.Method == "POST" {
		var c category.Category
		json.NewDecoder(w.Body).Decode(&c)
		resp := c.Add()
		if resp["status"] == "201" {
			fmt.Fprintf(r, "Created")
		} else {
			fmt.Fprintf(r, "Failed")
		}
	}
}

func CategoryDelete(r http.ResponseWriter, w *http.Request) {
	if w.Method == "DELETE" {
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
}

func CategoryUpdate(r http.ResponseWriter, w *http.Request) {
	if w.Method == "PATCH" {
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
}

func CategoryGetAllProjects(w http.ResponseWriter, r *http.Request) {

}
