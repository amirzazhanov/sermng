// Package httpserver to test server implementation
package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Record structure
type Record struct {
	ID          uint32 `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Counter     uint32 `json:"counter,omitempty"`
	URL         string `json:"url,omitempty"`
}

// Records defines slice of Record
type Records []Record

// RecordsStore main records store
var RecordsStore Records

// CreateRecord - Create new record in RecordStore
//=======================================================
func CreateRecord(w http.ResponseWriter, r *http.Request) {
	var rec Record
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	log.Println("BODY DUMP:", string(body))
	if err != nil {
		log.Fatalln("Error AddRecord", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddRecord", err)
	}
	if err := json.Unmarshal(body, &rec); err != nil { // unmarshall body contents
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error AddRecord unmarshalling data", err)
			return
		}
		return
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
		log.Println("=> data writen")
	}
	//		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}

// ReadRecord - Read specific record from RecordStore
//=======================================================
func ReadRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["record_id"], 10, 32)
	if err != nil {
		panic(err)
	}
	for i := range RecordsStore {
		if RecordsStore[i].ID == uint32(id) {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(RecordsStore[i]); err != nil {
				panic(err)
			}
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	return
}

// ReadAllRecords - http handler
//=======================================================
func ReadAllRecords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(RecordsStore); err != nil {
		panic(err)
	}
	return
}

// UpdateRecord - http handler
func UpdateRecord(w http.ResponseWriter, r *http.Request) {
	var rec Record
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["record_id"], 10, 32)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	log.Println("BODY DUMP:", string(body))
	if err != nil {
		log.Fatalln("Error Change Record", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error Change Record", err)
	}
	if err := json.Unmarshal(body, &rec); err != nil { // unmarshall body contents
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error Change Record unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}
	log.Println("new_description:", rec.Description, "new_counter:", rec.Counter, "new_url:", rec.URL)
	for i := range RecordsStore { // change Record
		if RecordsStore[i].ID == uint32(id) {
			if len(rec.Description) > 0 {
				RecordsStore[i].Description = rec.Description
			}
			RecordsStore[i].Counter = rec.Counter
			if len(rec.URL) > 0 {
				RecordsStore[i].URL = rec.URL
			}
			binBuffer, err := json.MarshalIndent(RecordsStore, "", "  ")
			if err != nil {
				log.Fatalln("Error Change Record marshalling data", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if err := ioutil.WriteFile(JSONFile.Name(), binBuffer, 0755); err != nil {
				log.Println("JSONFFile write:", err)
			} else {
				log.Println("=> data has been writen")
			}
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	return
}

// DeleteRecord - http handler
//=======================================================
func DeleteRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["record_id"], 10, 32)
	if err != nil {
		panic(err)
	}
	for i := range RecordsStore { // delete Record
		if RecordsStore[i].ID == uint32(id) {
			RecordsStore = append(RecordsStore[:i], RecordsStore[i+1:]...)
			binBuffer, err := json.MarshalIndent(RecordsStore, "", "  ")
			if err != nil {
				log.Fatalln("Error Delete Record marshalling data", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if err := ioutil.WriteFile(JSONFile.Name(), binBuffer, 0755); err != nil {
				log.Println("JSONFFile write:", err)
			} else {
				log.Println("=> data has been writen")
			}
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	return
}
