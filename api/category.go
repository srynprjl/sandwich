package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/srynprjl/sandwich/internal/logic"
)

func CategoryGetAll(r http.ResponseWriter, w *http.Request) {
	r.Header().Set("Content-Type", "application/json")
	resp := logic.GetAll()
	if resp["status"] == "200" {
		data := resp["data"].([]map[string]any)
		json.NewEncoder(r).Encode(data)
	} else {
		fmt.Fprintln(r, resp["message"].(string))
	}
}

func CategoryAdd(r http.ResponseWriter, w *http.Request) {
	var c logic.Category
	var data = make(map[string]any)
	json.NewDecoder(w.Body).Decode(&data)
	resp := c.Add(data)
	if resp["status"] == "201" {
		fmt.Fprintln(r, resp["message"])
	} else {
		fmt.Fprintln(r, resp["message"])
	}
}

func CategoryDelete(r http.ResponseWriter, w *http.Request) {
	id, err := strconv.Atoi(w.PathValue("id"))
	if err != nil {
		fmt.Fprint(r, err.Error())
	}
	c := logic.Category{Id: id}
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
	c := logic.Category{Id: id}
	updateData := make(map[string]any)
	json.NewDecoder(w.Body).Decode(&updateData)
	resp := c.Update(updateData)
	fmt.Fprint(r, resp["message"])

}

func CategoryGetAllProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		fmt.Fprintf(w, "Error")
		return
	}
	c := logic.Category{Id: id}
	resp := logic.GetProjects(c)
	if resp["status"] != "200" {
		fmt.Fprintln(w, resp["message"])
		return
	}
	json.NewEncoder(w).Encode(resp["data"].([]map[string]any))
}
