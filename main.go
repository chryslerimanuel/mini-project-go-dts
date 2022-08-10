package main

import (
	"fmt"
	"mini-project-go-dts/controllers"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Server started")

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	//root
	http.HandleFunc("/", controllers.HomeHandler)

	// user
	http.HandleFunc("/users", controllers.LoadDataUserHandler)
	http.HandleFunc("/users/index", controllers.LoadDataUserHandler)
	http.HandleFunc("/users/view", controllers.SelectOneUserHandler)
	http.HandleFunc("/users/add", controllers.CreateUserHandler)
	http.HandleFunc("/users/edit", controllers.UpdateUserHandler)
	http.HandleFunc("/users/delete", controllers.DeleteUserHandler)

	// task
	http.HandleFunc("/tasks", controllers.LoadDataTaskHandler)
	http.HandleFunc("/tasks/index", controllers.LoadDataTaskHandler)
	http.HandleFunc("/tasks/view", controllers.SelectOneTaskHandler)
	http.HandleFunc("/tasks/add", controllers.CreateTaskHandler)
	http.HandleFunc("/tasks/edit", controllers.UpdateTaskHandler)
	http.HandleFunc("/tasks/delete", controllers.DeleteTaskHandler)
	http.HandleFunc("/tasks/updatestatus", controllers.UpdateStatusHandler)

	fileServer := http.FileServer(http.Dir("assets"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	http.ListenAndServe(":"+port, nil)
}
