# GoWebApp

GoWebApp is a simple framework for writing web applications in Go.

You simply need to implement the URLHandler interface and register
a route to which that URLHandler responds. The simplest implementation
is the NullApp in the examples directory which just uses the DefaultHandler
to serve everything and respond with a 405 Method Not Allowed error.

```go
func main() {
	handle := URLHandler.DefaultHandler{}

	URLHandler.RegisterHandler(handle, "/")
	http.ListenAndServe(":8080", nil)
}
```

That's not terribly useful, so the SimpleGetApp example extends that
a little by implementing the Get method on a new type

```go
// Implement a new type which inherits the DefaultHandler
// behaviour
type SimpleGetPage struct {
	URLHandler.DefaultHandler
}

// Override Get to return some text for the root page, and
// a 404 error for any other request
func (r SimpleGetPage) Get(req *http.Request, params map[string]interface{}) (string, error) {
	if req.URL.Path == "/" {
		return "I am a page", nil
	}
	return "Page not found", URLHandler.NotFoundError{}

}

func main() {
	handle := SimpleGetPage{}

	URLHandler.RegisterHandler(handle, "/")
	http.ListenAndServe(":8080", nil)
}
```

Finally, the ListApp example in the examples folder shows a more complicated
application that implements both Get and Post to store/retrieve a list from 
an SQLite database and demonstrates the RegisterExtraParameter function
>>>>>>> Added README
