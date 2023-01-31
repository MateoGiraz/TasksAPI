package main

import (
	"github.com/MateoGiraz/TasksAPI/db"
	"github.com/MateoGiraz/TasksAPI/models"
	"github.com/MateoGiraz/TasksAPI/routes"
	"net/http"
)

func main() {
	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := routes.CreateRoutes()

	http.ListenAndServe(":3000", r)

}
