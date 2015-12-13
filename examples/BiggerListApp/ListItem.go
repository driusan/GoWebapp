package main

import (
	"bytes"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"github.com/driusan/GoWebapp/URLHandler"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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

	buf := new(bytes.Buffer)
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

// Implement an ETag handler for ListItems that just calculates an MD5 sum
// of the value.
//
// In a real application, this should be done on inserting/modifying
// the data and saved alongside the resource in the database, so that
// this just needs to read it from the DB, but since this is mostly
// implemented for the If-Match requirement of PUT and DELETE and this
// is just a demonstration application, the inefficiency (and poor
// choice of algorithm) should be acceptable..
func (r ListItem) ETag(url *url.URL, params map[string]interface{}) (URLHandler.ETag, error) {
	var db = params["database"].(*sql.DB)
	id := extractId(url)

	var value string
	err := db.QueryRow("SELECT Value FROM listitems WHERE ID=?", id).Scan(&value)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", URLHandler.NotFoundError{}
		}
		return "Unknown Error Querying Database", err
	}

	md5sum := md5.Sum([]byte(value))

	return URLHandler.ETag(hex.EncodeToString(md5sum[:])), nil
}
