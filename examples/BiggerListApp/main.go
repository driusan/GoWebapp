package main

import (
	"database/sql"
	"fmt"
	"github.com/driusan/GoWebapp/URLHandler"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

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
	URLHandler.RegisterExtraParameter("title", "I am a list")
	URLHandler.RegisterExtraParameter("database", db)

	// Register the ListPage as the handler for the root
	URLHandler.RegisterStaticHandler("/js/", "./js")
	URLHandler.RegisterHandler(ListItem{}, "/items/")
	URLHandler.RegisterHandler(ListPage{}, "/")
	fmt.Printf("Listening and serving on port 8080\n")

	// Let Go's http library handle the heavy lifting
	http.ListenAndServe(":8080", nil)
}
