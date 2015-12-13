package main

import (
	"github.com/driusan/GoWebapp/URLHandler"
	"database/sql"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"bytes"
)

func extractId(URL *url.URL) int {
	urlComponents := strings.Split(URL.Path, "/")
	i, err := strconv.Atoi(urlComponents[2])
	if err != nil {
		panic("Couldn't extract id")
	}
	return i
}

type ListItem struct {
	URLHandler.DefaultHandler
}

func (r ListItem) Get(req *http.Request, extraparams map[string]interface{}) (string, error) {
	db := extraparams["database"].(*sql.DB)
	id := extractId(req.URL)

	type Item struct {
		ID    int
		Value string
	}
	// Populate the pageData items from the database
	var value string
	err := db.QueryRow("SELECT Value FROM listitems WHERE ID=?", id).Scan(&value)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", URLHandler.NotFoundError{}
		}
		return "Unknown Error Querying Database", err
	}
	return value, nil
}

func (r ListItem) Delete(req *http.Request, extraparams map[string]interface{}) (string, error) {
	var db = extraparams["database"].(*sql.DB)
	id := extractId(req.URL)
	res, err := db.Exec("DELETE FROM listitems WHERE ID=?", id)
	if err != nil {
		return "Unknown error", err
	}

	aff, err := res.RowsAffected()
	if err != nil {
		return "Unknown error", err
	}

	if aff <= 0 {
		return "", URLHandler.NotFoundError{}
	}
	return "Deleted", nil
}

func (r ListItem) Put(req *http.Request, extraparams map[string]interface{}) (string, error) {
	var db = extraparams["database"].(*sql.DB)
	id := extractId(req.URL)

	buf := new (bytes.Buffer)
	buf.ReadFrom(req.Body)

	value := buf.String()

	res, err := db.Exec("UPDATE listitems SET Value=? where id=?", value, id)
	if err != nil {
		return "Unknown error", err
	}

	aff, err := res.RowsAffected()
	if err != nil {
		return "Unknown error", err
	}

	if aff <= 0 {
		return "", URLHandler.NotFoundError{}
	}
	return "Updated", nil
}
func (r ListItem) ETag(url *url.URL, params map[string]interface{}) URLHandler.ETag {
	return ""
}
