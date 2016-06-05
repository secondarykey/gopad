package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func init() {

	http.HandleFunc("/", listHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)

	http.HandleFunc("/test", createHandler)
	// presentation

}

func main() {
	err := listen("gopad.db")
	if err != nil {
		fmt.Printf("gopad listen Error %v\n", err)
		os.Exit(1)
	}
}

func listen(file string) error {

	_, err := os.Stat(file)
	flag := err == nil

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return fmt.Errorf("Database Open Error : %v", err)
	}
	fmt.Println("###### Serve Database")

	if !flag {
		_, err := db.Exec("CREATE TABLE memos(ID INTEGER PRIMARY KEY AUTOINCREMENT,TITLE VARCHAR(255),CONTENT TEXT)")
		if err != nil {
			return fmt.Errorf("Create Table Error : %v", err)
		}
	}

	Use(db)

	http.Handle("/static/", http.FileServer(http.Dir("")))

	fmt.Println("###### Serve Web")
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

	fmt.Println("List Handler")

	tc := make(map[string]interface{})

	memoList, err := Memo{}.All().Query()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tc["MemoList"] = memoList

	setTemplates(w, tc, "list.tmpl")
}

func createHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Create Handler")

	memo := Memo{
		Title:   "default",
		Content: "default",
	}

	_, err := memo.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	key := fmt.Sprintf("%d", memo.Id)
	fmt.Println(key)

	http.Redirect(w, r, "/view/"+key, 301)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("View Handler")

	url := r.URL.Path
	pathS := strings.Split(url, "/")
	key := pathS[2]

	id, err := strconv.Atoi(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	memo, err := Memo{}.Find(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	tc := make(map[string]interface{})
	tc["Memo"] = memo

	setTemplates(w, tc, "view.tmpl")
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	pathS := strings.Split(url, "/")
	key := pathS[2]
	memo, _ := Memo{}.FindBy("id", key)

	tc := make(map[string]interface{})
	tc["Memo"] = memo

	setTemplates(w, tc, "edit.tmpl")
}
