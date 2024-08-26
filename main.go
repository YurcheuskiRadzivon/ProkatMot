package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home_page.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.ExecuteTemplate(w, "home_page", nil)
}
func recording(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/recording.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.ExecuteTemplate(w, "recording", nil)
}

var database *sql.DB

func saveArticle(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	phonenum := r.FormValue("phonenum")

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/forms")
	if err != nil {
		panic(err)

	}
	defer db.Close()
	insert, err := db.Query(fmt.Sprintf("INSERT INTO `users` (`name`,`phonenumber`) VALUES('%s','%s')", name, phonenum))
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
func Handler() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", homePage)
	http.HandleFunc("/recording/", recording)
	http.HandleFunc("/save_article", saveArticle)

	http.ListenAndServe(":8080", nil)
}
func main() {
	Handler()
}
