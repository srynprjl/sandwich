package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/srynprjl/sandwich/internal/projects"
)

func ProjectGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	catId, catErr := strconv.Atoi(r.PathValue("catId"))
	if err != nil || catErr != nil {
		fmt.Fprintf(w, "ERROR")
		return
	}
	project := projects.Project{Id: id, Category: catId}
	resp := project.Get()
	if resp["status"] != "200" {
		fmt.Fprintln(w, resp["message"].(string))
		return
	}
	json.NewEncoder(w).Encode(resp["data"].(projects.Project))
}

func ProjectGetRandom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := projects.GetRandom()
	if resp["status"] != "200" {
		fmt.Fprintln(w, resp["message"].(string))
		return
	}
	json.NewEncoder(w).Encode(resp["data"].(projects.Project))
}

func ProjectGetNRandom(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.Atoi(r.PathValue("num"))
	if err != nil {
		fmt.Fprintln(w, "ERROR")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	resp := projects.GetNRandom(n)
	if resp["status"] != "200" {
		fmt.Fprintln(w, resp["message"].(string))
		return
	}
	json.NewEncoder(w).Encode(resp["data"].([]projects.Project))
}

func ProjectAdd(w http.ResponseWriter, r *http.Request) {
	catId, err := strconv.Atoi(r.PathValue("catId"))
	if err != nil {
		fmt.Fprintln(w, "Error")
		return
	}
	p := projects.Project{Category: catId}
	json.NewDecoder(r.Body).Decode(&p)
	resp := p.Add()
	fmt.Fprintln(w, resp["message"].(string))
}

func ProjectDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	catId, catErr := strconv.Atoi(r.PathValue("catId"))
	if err != nil || catErr != nil {
		fmt.Fprintln(w, "Error")
		return
	}
	p := projects.Project{Id: id, Category: catId}
	resp := p.Remove()
	fmt.Fprintln(w, resp["message"].(string))
}

func ProjectUpdate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	catId, catErr := strconv.Atoi(r.PathValue("catId"))
	if err != nil || catErr != nil {
		fmt.Fprintln(w, "Error")
		return
	}
	p := projects.Project{Id: id, Category: catId}
	var data map[string]any
	jsonErr := json.NewDecoder(r.Body).Decode(&data)
	if jsonErr != nil {
		fmt.Fprintln(w, "Error!")
	}
	fmt.Println(p, data)
	resp := p.Update(data)
	fmt.Fprintln(w, resp["message"])
}
