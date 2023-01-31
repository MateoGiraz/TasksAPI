package routes

import (
	"github.com/gorilla/mux"
)

func CreateRoutes() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)

	r.HandleFunc("/users", GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", GetUserHandler).Methods("GET")
	r.HandleFunc("/users", PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", DeleteUserHandler).Methods("DELETE")

	//task routes

	r.HandleFunc("/tasks", GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks", CreateTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", DeleteTaskHandler).Methods("DELETE")

	return r
}
