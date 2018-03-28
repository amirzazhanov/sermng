// Package httpserver to test server implementation
package main

import (
	"html/template"
	"net/http"
)

// HandlerHTML - handler for /rss/ url
func HandlerHTML(w http.ResponseWriter, r *http.Request) {
	pTmp := &Page{Title: "TestPage", Body: []byte("This is a sample simple Page.")}
	t, _ := template.ParseFiles("template.tpl", "bootstrap4_css.tpl")
	t.Execute(w, pTmp)
}

// HandlerRecords ...
func HandlerRecords(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
	} else if r.Method == "PUT" {
	} else if r.Method == "DELETE" {
	} else if r.Method == "GET" {
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
