// Package sermng to test server implementation
package main

// https://golang.org/doc/articles/wiki/
// http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/
// https://thenewstack.io/make-a-restful-json-api-go/

import (
	"log"
	"net/http"
)

// Page structure simple
type Page struct {
	Title string
	Body  []byte
}

// Log function
func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s\t%s\t%s\t%s\tX-FORWARDED-FOR: %s", r.RemoteAddr, r.Method, r.Proto, r.URL.Path[1:], r.Header.Get("X-FORWARDED-FOR"))
		handler.ServeHTTP(w, r)
	})
}

func main() {
	http.HandleFunc("/html/", HandlerHTML)
	http.HandleFunc("v1/records", HandlerRecords)
	log.Fatal(http.ListenAndServe(":8080", Log(http.DefaultServeMux)))
}
