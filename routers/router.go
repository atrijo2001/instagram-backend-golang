package routers

import (
	"github.com/arrijo2001/insta/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", controllers.GetAUser).Methods("GET")

	return r
}
