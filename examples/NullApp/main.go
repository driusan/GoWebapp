package main

import (
	"github.com/driusan/GoWebapp/URLHandler"
	"net/http"
)

func main() {
	handle := URLHandler.DefaultHandler{}

	URLHandler.RegisterHandler(handle, "/")
	http.ListenAndServe(":8080", nil)
}
