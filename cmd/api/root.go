package api

import (
	"fmt"
	"net/http"
)

func Api() {
	mux := http.NewServeMux()

	// categories

	mux.HandleFunc("GET /api/category/", CategoryGetAll)
	mux.HandleFunc("POST /api/category/", CategoryAdd)
	mux.HandleFunc("PATCH /api/category/{id}/", CategoryUpdate)
	mux.HandleFunc("DELETE /api/category/{id}/", CategoryDelete)
	mux.HandleFunc("GET /api/category/{id}", CategoryGetAllProjects)

	// projects
	mux.HandleFunc("GET /api/projects/{id}/", ProjectGet)
	mux.HandleFunc("GET /api/projects/random/", ProjectGetRandom)
	mux.HandleFunc("GET /api/projects/random/{num}", ProjectGetNRandom)
	mux.HandleFunc("POST /api/projects/", ProjectAdd)
	mux.HandleFunc("PATCH /api/projects/{id}/", ProjectUpdate)
	mux.HandleFunc("DELETE /api/projects/{id}/", ProjectDelete)

	// server start
	fmt.Println("Starting server...!")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("ERROR!")
	}

}
