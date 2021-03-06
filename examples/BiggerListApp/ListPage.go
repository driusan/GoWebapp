package main

import (
	"bytes"
	"database/sql"
	"github.com/driusan/GoWebapp/URLHandler"
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
