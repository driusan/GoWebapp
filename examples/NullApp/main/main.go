package main

import (
	"net/http"
	"URLHandler"
)

func main() {
	handle := URLHandler.DefaultHandler{}

	URLHandler.RegisterHandler(handle, "/")
	http.ListenAndServe(":8080", nil)
}
