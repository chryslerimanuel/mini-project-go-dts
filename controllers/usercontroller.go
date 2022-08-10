package controllers

import (
	"mini-project-go-dts/entities"
	"log"
	"net/http"
	"path"
	"strconv"
	"text/template"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		log.Println("Method GET dipakai.")

		tmpl, err := template.ParseFiles(path.Join("views", "user", "create.html"), path.Join("views", "layout.html"))
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

	name := r.FormValue("name")
	roleId := r.FormValue("roleId")

	roleIdInt, err := strconv.Atoi(roleId)
	if err != nil || roleIdInt < 1 {
		log.Printf("Errors %s converting roleId to integer", err)
		panic(err.Error())
	}

	user := entities.User{Name: name, RoleId: roleIdInt}

	err = user.Insert()
	if err != nil {
		panic(err.Error())
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	IdInt, err := strconv.Atoi(id)
	if err != nil || IdInt < 1 {
		log.Printf("Errors %s converting to integer", err)
		panic(err.Error())
	}

	var user entities.User
	err = user.Delete(IdInt)
	if err != nil {
		panic(err.Error())
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	name := r.FormValue("name")
	roleId := r.FormValue("roleId")

	idInt, err := strconv.Atoi(id)
	if err != nil || idInt < 1 {
		log.Printf("Errors %s converting to integer", err)
		panic(err.Error())
	}

	roleIdInt, err := strconv.Atoi(roleId)
	if err != nil || roleIdInt < 1 {
		log.Printf("Errors %s converting to integer", err)
		panic(err.Error())
	}

	user := entities.User{Name: name, RoleId: roleIdInt}

	err = user.Update(idInt)
	if err != nil {
		panic(err.Error())
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func LoadDataUserHandler(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	res, err := user.FindAll()
	if err != nil {
		panic(err.Error())
	}

	tmpl, err := template.ParseFiles(path.Join("views", "user", "loaddatagrid.html"), path.Join("views", "layout.html"))
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

func SelectOneUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	state := r.URL.Query().Get("state")

	idInt, err := strconv.Atoi(id)
	if err != nil || idInt < 1 {
		log.Printf("Errors %s converting id to integer", err)
		http.NotFound(w, r)
		return
	}

	var user entities.User
	res, err := user.Select(idInt)
	if err != nil {
		panic(err.Error())
	}

	if state == "2" {
		tmpl, err := template.ParseFiles(path.Join("views", "user", "update.html"), path.Join("views", "layout.html"))
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
		tmpl, err := template.ParseFiles(path.Join("views", "user", "view.html"), path.Join("views", "layout.html"))
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
