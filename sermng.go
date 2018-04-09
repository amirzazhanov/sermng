// Package sermng to test server implementation
package main

// https://golang.org/doc/articles/wiki/
// http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/
// https://thenewstack.io/make-a-restful-json-api-go/

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// JSONFile GLOBAL FILE pointer to main JSON FILE
var JSONFile *os.File

// Log function
func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s\t%s\t%s\t%s\tX-FORWARDED-FOR: %s", r.RemoteAddr, r.Method, r.Proto, r.URL.Path[1:], r.Header.Get("X-FORWARDED-FOR"))
		handler.ServeHTTP(w, r)
	})
}

func main() {
	// Open our jsonFile
	var err error
	JSONFile, err = os.OpenFile("records.json", os.O_RDWR, 0755)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer JSONFile.Close()
	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(JSONFile)
	// we unmarshal our byteValue which contains our
	// jsonFile's content into 'RecordsStore' which we defined above
	json.Unmarshal(byteValue, &RecordsStore)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/v1/records", HandlerRecords)
	http.HandleFunc("/v1/records/", HandlerRecords)
	log.Fatal(http.ListenAndServe(":8080", Log(http.DefaultServeMux)))
}
