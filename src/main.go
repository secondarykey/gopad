package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pborman/uuid"
)

func init() {

	http.HandleFunc("/", listHandler)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)

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

	http.Handle("/static/", http.FileServer(http.Dir("")))
	return http.ListenAndServe(":5005", nil)
}

func setTemplates(w http.ResponseWriter, p interface{}, files ...string) {

	templateDir := "templates"

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

	memoList := Memo{}.Query()
	tc["MemoList"] = memoList

	setTemplates(w, tc, "list.tmpl")
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	id := uuid.New()
	memo := Memo{Id: id}

	memo.Save()
	http.Redirect(w, r, "/view/"+id, 301)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	tc := make(map[string]interface{})

	setTemplates(w, tc, "view.tmpl")
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	tc := make(map[string]interface{})

	setTemplates(w, tc, "edit.tmpl")
}
