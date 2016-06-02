package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"net/http"
	"os"
)

func init() {

	// list
	http.HandleFunc("/", listHandler)

	// view

	// edit

	// presentation

}

var db *sql.DB

func main() {
	err := listen()
	if err != nil {
		fmt.Printf("gopad listen Error %v\n", err)
		os.Exit(1)
	}
}

func listen() error {

	_, err := sql.Open("sqlite3", "gopad.db")
	if err != nil {
		return fmt.Errorf("Database Open Error : %v", err)
	}

	http.Handle("/static/", http.FileServer(http.Dir("..")))
	return http.ListenAndServe(":5005", nil)
}

func setTemplates(w http.ResponseWriter, p interface{}, files ...string) {

	templateDir := "../templates"

	tmpls := make([]string, 0)
	tmpls = append(tmpls, templateDir+"/layout.tmpl")

	for _, elm := range files {
		tmpls = append(tmpls, templateDir+"/"+elm)
	}

	tmpl := template.Must(template.ParseFiles(tmpls...))
	if err := tmpl.Execute(w, p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func listHandler(w http.ResponseWriter, r *http.Request) {

	tc := make(map[string]interface{})
	tc["MemoList"] = nil
	tc["User"] = nil

	setTemplates(w, tc, "list.tmpl")
}
