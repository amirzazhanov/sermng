// Package httpserver to test server implementation
package main

import (
	"encoding/json"
	"html/template"
	"net/http"
)

// Record structure
type Record struct {
	ID          uint32 `json:"id"`
	Description string `json:"description"`
	Counter     uint32 `json:"counter"`
	URL         string `json:"url"`
}

// Records defines slice of Record
type Records []Record

// init main RecordsStore variable
var RecordsStore Records

// HandlerHTML - handler for /rss/ url
func HandlerHTML(w http.ResponseWriter, r *http.Request) {
	pTmp := &Page{Title: "TestPage", Body: []byte("This is a sample simple Page.")}
	t, _ := template.ParseFiles("template.tpl", "bootstrap4_css.tpl")
	t.Execute(w, pTmp)
}

// HandlerRecords ...
func HandlerRecords(w http.ResponseWriter, r *http.Request) {
	//	records := Records{
	//		Record{ID: 1, Description: "test123", Counter: 15, URL: "http://test123.com"},
	//		Record{ID: 2, Description: "test456", Counter: 26, URL: "http://test456.com"},
	//	}
	if r.Method == "POST" {
	} else if r.Method == "PUT" {
	} else if r.Method == "DELETE" {
	} else if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(RecordsStore); err != nil {
			panic(err)
		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
