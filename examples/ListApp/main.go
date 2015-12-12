package main

import (
	"github.com/driusan/GoWebapp/URLHandler"
	"bytes"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"net/http"
)

type ListPage struct {
	URLHandler.DefaultHandler
}

func (r ListPage) Get(req *http.Request, extraparams map[string]interface{}) (string, error) {
	db := extraparams["database"].(*sql.DB)
	title := extraparams["title"].(string)

	type Item struct {
		ID    int
		Value string
	}
	pageData := struct {
		Title string
		Items []Item
	}{Title: title}

	// Populate the pageData items from the database
	rows, err := db.Query("SELECT ID, Value FROM listitems")
	defer rows.Close()
	if err != nil {
		return "Error Querying Database", err
	}

	for rows.Next() {
		var value string
		var ID int
		rows.Scan(&ID, &value)
		i := Item{ID, value}
		pageData.Items = append(pageData.Items, i)
	}

	// Now render the page template
	pageTemplate, err := template.ParseFiles("templates/main.html")
	if err != nil {
		panic("Couldn't parse template file")
	}

	// Render the template to a bytes Buffer, so that we can return
	// the rendered string. We don't have access to the ResponseWriter
	// here.
	pageBuffer := new(bytes.Buffer)
	err = pageTemplate.Execute(pageBuffer, pageData)
	if err != nil {
		panic("Could not execute main template")
	}
	return pageBuffer.String(), nil
}

func (r ListPage) Post(req *http.Request, extraparams map[string]interface{}) (string, string, error) {
	var db = extraparams["database"].(*sql.DB)
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}
	insert, err := db.Prepare("INSERT INTO listitems (Value) VALUES (?)")
	if err != nil {
		panic(err)
	}
	_, err = insert.Exec(req.PostForm.Get("input"))
	if err != nil {
		panic(err)
	}
	return "Data has been inserted", "/", nil
}

func main() {
	// Initialize an sqlite3 database connection
	db, err := sql.Open("sqlite3", "list.db")
	defer db.Close()
	if err != nil {
		panic("Couldn't open db")
	}

	// Create a ListPage instance, and register a title
	// (for fun), and a reference to the database object
	// (so that our implementation can interact with the DB)
	x := ListPage{}
	URLHandler.RegisterExtraParameter("title", "I am a list")
	URLHandler.RegisterExtraParameter("database", db)

	// Register the ListPage as the handler for the root
	URLHandler.RegisterHandler(x, "/")
	fmt.Printf("Listening and serving on port 8080\n")

	// Let Go's http library handle the heavy lifting
	http.ListenAndServe(":8080", nil)
}
