// Package httpserver to test server implementation
package main

import (
	"encoding/json"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Record structure
type Record struct {
	ID          uint32 `json:"id,omitempty"`
	Description string `json:"description"`
	Counter     uint32 `json:"counter"`
	URL         string `json:"url"`
}

// Records defines slice of Record
type Records []Record

// RecordsStore main records store
var RecordsStore Records

// HandlerHTML - handler for /rss/ url
func HandlerHTML(w http.ResponseWriter, r *http.Request) {
	pTmp := &Page{Title: "TestPage", Body: []byte("This is a sample simple Page.")}
	t, _ := template.ParseFiles("template.tpl")
	t.Execute(w, pTmp)
}

// HandlerRecords ...
func HandlerRecords(w http.ResponseWriter, r *http.Request) {
	var rec Record
	// --------------------------------------------------------
	// --------------------- POST HANDLER ---------------------
	// --------------------------------------------------------
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
		log.Println(body)
		if err != nil {
			log.Fatalln("Error AddRecord", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := r.Body.Close(); err != nil {
			log.Fatalln("Error AddRecord", err)
		}
		if err := json.Unmarshal(body, &rec); err != nil { // unmarshall body contents
			w.WriteHeader(422) // unprocessable entity
			log.Println(err)
			if err := json.NewEncoder(w).Encode(err); err != nil {
				log.Fatalln("Error AddRecord unmarshalling data", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		log.Println(rec)
		for _, recstor := range RecordsStore { // autoincrement ID
			if recstor.ID >= rec.ID {
				rec.ID = recstor.ID + 1
			}
		}
		RecordsStore = append(RecordsStore, rec)
		binBuffer, err := json.MarshalIndent(RecordsStore, "", "  ")
		if err != nil {
			log.Fatalln("Error AddRecord marshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := ioutil.WriteFile(JSONFile.Name(), binBuffer, 0755); err != nil {
			log.Println("JSONFFile write:", err)
		} else {
			log.Println("==>>> data writen")
		}
		//		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		return
		// -------------------------------------------------------
		// --------------------- PUT HANDLER ---------------------
		// -------------------------------------------------------
	} else if r.Method == "PUT" {
		id := getID(r, 0)
		var putSuccess bool
		log.Println("PUT id =", id)
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
		log.Println(body)
		if err != nil {
			log.Fatalln("Error Change Record", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := r.Body.Close(); err != nil {
			log.Fatalln("Error Change Record", err)
		}
		if err := json.Unmarshal(body, &rec); err != nil { // unmarshall body contents
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			if err := json.NewEncoder(w).Encode(err); err != nil {
				log.Fatalln("Error Change Record unmarshalling data", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		log.Println(rec)
		for i := range RecordsStore { // change Record
			if RecordsStore[i].ID == id {
				RecordsStore[i].Description = rec.Description
				RecordsStore[i].Counter = rec.Counter
				RecordsStore[i].URL = rec.URL
				putSuccess = true
				break
			}
		}
		if putSuccess {
			binBuffer, err := json.MarshalIndent(RecordsStore, "", "  ")
			if err != nil {
				log.Fatalln("Error Change Record marshalling data", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if err := ioutil.WriteFile(JSONFile.Name(), binBuffer, 0755); err != nil {
				log.Println("JSONFFile write:", err)
			} else {
				log.Println("==>>> data writen")
			}
			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusNotFound)
		return
		// ----------------------------------------------------------
		// --------------------- DELETE HANDLER ---------------------
		// ----------------------------------------------------------
	} else if r.Method == "DELETE" {
		var deleteSuccess bool // need to remove this var, but not now
		id := getID(r, 0)
		for i := range RecordsStore { // delete Record
			if RecordsStore[i].ID == id {
				RecordsStore = append(RecordsStore[:i], RecordsStore[i+1:]...)
				deleteSuccess = true
				break
			}
		}
		if deleteSuccess {
			binBuffer, err := json.MarshalIndent(RecordsStore, "", "  ")
			if err != nil {
				log.Fatalln("Error Delete Record marshalling data", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if err := ioutil.WriteFile(JSONFile.Name(), binBuffer, 0755); err != nil {
				log.Println("JSONFFile write:", err)
			} else {
				log.Println("==>>> data writen")
			}
			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusNotFound)
		return
		// -------------------------------------------------------
		// --------------------- GET HANDLER ---------------------
		// -------------------------------------------------------
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

// Extract a code from a URL. Return the default code if code
// is missing or code is not a valid number.
func getID(r *http.Request, defaultID uint32) uint32 {
	p := strings.Split(r.URL.Path, "/v1/records/")
	if len(p) == 1 {
		return defaultID
	} else if len(p) > 1 {
		id, err := strconv.ParseUint(p[1], 10, 32)
		if err == nil {
			return uint32(id)
		}
		return defaultID
	} else {
		return defaultID
	}
}
