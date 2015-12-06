package main

import (
	"URLHandler"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

type ListPage struct {
	URLHandler.DefaultHandler
}

func (r ListPage) Get(req *http.Request, extraparams map[string]interface{}) (string, error) {
	db := extraparams["database"].(*sql.DB)
	template := "<html><body>"
	template += "<h3>" + extraparams["title"].(string) + "</h3>"
	template += "	<form method=\"post\"><ul>"

	rows, err := db.Query("SELECT Value FROM listitems")
	defer rows.Close()
	if err != nil {
		return "Error Querying Database", err
	}

	for rows.Next() {
		var value string
		rows.Scan(&value)
		template += "<li>" + value + "</li>"
	}
	template += "</ul><input type=\"text\" name=\"input\"/><input type=\"submit\" /></form></body></html>"

	return template, nil
}

func (r ListPage) Post(req *http.Request, extraparams map[string]interface{}) (string, string, error) {
	var db = extraparams["database"].(*sql.DB)
	err := req.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	insert, err := db.Prepare("INSERT INTO listitems (Value) VALUES (?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = insert.Exec(req.PostForm.Get("input"))
	if err != nil {
		log.Fatal(err)
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
