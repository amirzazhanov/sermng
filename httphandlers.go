// Package httpserver to test server implementation
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// HandlerImages - handler for /images/ url
func HandlerImages(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Requested Image Is %s!", r.URL.Path[1:])
}

// HandlerHTML - handler for /rss/ url
func HandlerHTML(w http.ResponseWriter, r *http.Request) {
	pTmp := &Page{Title: "TestPage", Body: []byte("This is a sample simple Page.")}
	t, _ := template.ParseFiles("template.html")
	t.Execute(w, pTmp)
}
