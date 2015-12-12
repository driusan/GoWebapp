package main

import (
	"github.com/driusan/GoWebapp/URLHandler"
	"net/http"
)

func main() {
	URLHandler.RegisterStaticHandler("/", "./static")
	http.ListenAndServe(":8080", nil)
}
