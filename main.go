package main

import (
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	// Serve static files from "static" directory.
	http.Handle("/", http.FileServer(http.Dir("public")))
	appengine.Main()
}
