package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"test-mux/database"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"test-mux/models"
)

var (
	DB *gorm.DB
)

type Users struct {
	Users []models.User
}

func main() {

	DB = database.Connection()

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/about", about).Methods("GET")
	r.HandleFunc("/create-user", createUser).Methods("POST")
	r.HandleFunc("/delete", deleteUser).Methods("GET")
	r.HandleFunc("/update", updateUser).Methods("POST")
	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	r.PathPrefix("/static/").Handler(s)

	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	//json.NewEncoder(w).Encode(map[string]bool{"detail": true})
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	var users []models.User
	DB.Find(&users)
	data := Users{Users: users}

	tmpl.Execute(w, data)
}

func createUser(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	email := r.FormValue("email")
	pass := r.FormValue("password")

	DB.Create(&models.User{Name: name, Email: email, Password: pass})

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	modelUser := new(models.User)
	DB.Delete(&modelUser, id)
	fmt.Println(id)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	name := r.FormValue("name")
	email := r.FormValue("email")
	pass := r.FormValue("password")
	modelUser := new(models.User)

	DB.Model(&modelUser).Where("ID = ?", id).Updates(map[string]interface{}{"Name": name, "Email": email, "Password": pass})
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func about(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/about.html"))
	tmpl.Execute(w, "data")
}
