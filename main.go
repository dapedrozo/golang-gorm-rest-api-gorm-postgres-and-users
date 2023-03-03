package main

import (
	"net/http"

	"github.com/dapedrozo/go-gorm-resapi/db"
	"github.com/dapedrozo/go-gorm-resapi/models"
	"github.com/dapedrozo/go-gorm-resapi/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")
	r.HandleFunc("/users/permanently/{id}", routes.DeleteUserPermanentlyHandler).Methods("DELETE")

	//Tasks routes
	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
