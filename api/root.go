package api

import (
	"fmt"
	"net/http"
)

func Api(host string, port int) {
	mux := http.NewServeMux()
	fmt.Println(host, port)
	mux.HandleFunc("GET /api/category", CategoryGetAll)
	mux.HandleFunc("POST /api/category", CategoryAdd)
	mux.HandleFunc("PATCH /api/category/{id}", CategoryUpdate)
	mux.HandleFunc("DELETE /api/category/{id}", CategoryDelete)
	mux.HandleFunc("GET /api/category/{id}", CategoryGetAllProjects)

	// projects
	mux.HandleFunc("GET /api/category/{catId}/projects/{id}", ProjectGet)
	mux.HandleFunc("GET /api/projects/random", ProjectGetRandom)
	mux.HandleFunc("GET /api/projects/random/{num}", ProjectGetNRandom)
	mux.HandleFunc("POST /api/category/{catId}/projects", ProjectAdd)
	mux.HandleFunc("PATCH /api/category/{catId}/projects/{id}", ProjectUpdate)
	mux.HandleFunc("DELETE /api/category/{catId}/projects/{id}", ProjectDelete)

	// server start
	fmt.Printf("Starting server...! at http://%s:%d", host, port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), mux)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("ERROR!")
	}

}
