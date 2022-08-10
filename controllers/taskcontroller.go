package controllers

import (
	"mini-project-go-dts/entities"
	"log"
	"net/http"
	"path"
	"strconv"
	"text/template"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		log.Println("Method GET dipakai.")

		tmpl, err := template.ParseFiles(path.Join("views", "task", "create.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Printf("Errors %s load template", err)
			http.Error(w, "An error has occured.", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Printf("Errors %s execute template", err)
			http.Error(w, "An error has occured.", http.StatusInternalServerError)
			return
		}

		return
	}

	taskDetail := r.FormValue("taskDetail")
	sendToName := r.FormValue("sendToName")
	taskDeadLine := r.FormValue("taskDeadLine")

	task := entities.Task{TaskDetail: taskDetail, CreatedByName: "Admin", SendToName: sendToName, TaskDeadLine: taskDeadLine}

	err := task.Insert()
	if err != nil {
		// panic(err.Error())
		http.Redirect(w, r, "/tasks/add", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/tasks", http.StatusSeeOther)
	}
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	IdInt, err := strconv.Atoi(id)
	if err != nil || IdInt < 1 {
		log.Printf("Errors %s converting to integer", err)
		panic(err.Error())
	}

	var task entities.Task
	err = task.Delete(IdInt)
	if err != nil {
		panic(err.Error())
	}

	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	taskDetail := r.FormValue("taskDetail")
	sendToName := r.FormValue("sendToName")
	taskDeadLine := r.FormValue("taskDeadLine")

	idInt, err := strconv.Atoi(id)
	if err != nil || idInt < 1 {
		log.Printf("Errors %s converting to integer", err)
		panic(err.Error())
	}

	task := entities.Task{TaskDetail: taskDetail, SendToName: sendToName, TaskDeadLine: taskDeadLine}

	err = task.Update(idInt)
	if err != nil {
		panic(err.Error())
	}

	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}

func LoadDataTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task entities.Task
	res, err := task.FindAll()
	if err != nil {
		panic(err.Error())
	}

	tmpl, err := template.ParseFiles(path.Join("views", "task", "loaddatagrid.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Printf("Errors %s load template", err)
		http.Error(w, "An error has occured.", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, res)
	if err != nil {
		log.Printf("Errors %s execute template", err)
		http.Error(w, "An error has occured.", http.StatusInternalServerError)
		return
	}
}

func SelectOneTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	state := r.URL.Query().Get("state")

	idInt, err := strconv.Atoi(id)
	if err != nil || idInt < 1 {
		log.Printf("Errors %s converting id to integer", err)
		http.NotFound(w, r)
		return
	}

	var task entities.Task
	res, err := task.Select(idInt)
	if err != nil {
		panic(err.Error())
	}

	if state == "2" {
		tmpl, err := template.ParseFiles(path.Join("views", "task", "update.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Printf("Errors %s load template", err)
			http.Error(w, "An error has occured.", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, res)
		if err != nil {
			log.Printf("Errors %s execute template", err)
			http.Error(w, "An error has occured.", http.StatusInternalServerError)
			return
		}

	} else {
		tmpl, err := template.ParseFiles(path.Join("views", "task", "view.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Printf("Errors %s load template", err)
			http.Error(w, "An error has occured.", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, res)
		if err != nil {
			log.Printf("Errors %s execute template", err)
			http.Error(w, "An error has occured.", http.StatusInternalServerError)
			return
		}
	}
}

func UpdateStatusHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	idInt, err := strconv.Atoi(id)
	if err != nil || idInt < 1 {
		log.Printf("Errors %s converting to integer", err)
		panic(err.Error())
	}

	task := entities.Task{}

	err = task.UpdateStatus(idInt)
	if err != nil {
		panic(err.Error())
	}

	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}
