# ListApp

This is a demonstration of a simple sqlite backed list application.

A GET request will show all items in the sqlite database, a POST
request will add to it and respond with a 303 See Other.

You'll have to:

```bash
go get github.com/mattn/go-sqlite3
```

before you can use this 

TODO: Write a ListItem handler which can deal with PUT and DELETE
