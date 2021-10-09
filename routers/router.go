package routers

import (
	"github.com/arrijo2001/insta/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", controllers.GetAUser).Methods("GET")

	r.HandleFunc("/posts", controllers.GetAllPosts).Methods("GET")
	r.HandleFunc("/posts", controllers.CreatePost).Methods("POST")
	r.HandleFunc("/posts/{id}", controllers.GetAPost).Methods("GET")

	return r
}
