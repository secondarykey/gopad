package main

import (
	"database/sql"
	"encoding/json"
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

	http.HandleFunc("/create", createHandler)
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

	memo := Memo{
		Title:   "Title",
		Content: "# Menu",
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

	m, err := getMemo(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tc := make(map[string]interface{})
	tc["Memo"] = m
	setTemplates(w, tc, "view.tmpl")
}

func getMemo(r *http.Request) (*Memo, error) {

	url := r.URL.Path
	pathS := strings.Split(url, "/")
	key := pathS[2]

	id, err := strconv.Atoi(key)
	if err != nil {
		return nil, err
	}

	m, err := Memo{}.Find(id)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func editHandler(w http.ResponseWriter, r *http.Request) {

	m, err := getMemo(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == "POST" {
		r.ParseForm()

		m.Title = r.FormValue("title")
		m.Content = r.FormValue("content")

		m.Save()

	} else if r.Method == "DELETE" {
		m.Destroy()
	} else {

		tc := make(map[string]interface{})
		tc["Memo"] = m

		setTemplates(w, tc, "edit.tmpl")
		return
	}

	w.WriteHeader(200)
	enc := json.NewEncoder(w)
	d := map[string]interface{}{
		"success": true,
	}
	enc.Encode(d)

}
