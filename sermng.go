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

// ConfigFile Global File pounter to configuration file
var ConfigFile *os.File

// ConfigFileName List of default locations for configuration filenames
var ConfigFileName = []string{"sermngconfig.json", "/etc/sermngconfig.json", "/usr/local/etc/sermngconfig.json"}

type Conf struct {
	RestAPIPort        uint32 `json:"restAPIPort,omitempty"`
	RestAPIBindAddress string `json:"restAPIBindAddress,omitempty"`
	DataJSONFile       string `json:"dataJSONFile,omitempty"`
}
type Configuration []Conf

var CFG Configuration

// Log function
func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s\t%s\t%s\t%s\tX-FORWARDED-FOR: %s", r.RemoteAddr, r.Method, r.Proto, r.URL.Path[1:], r.Header.Get("X-FORWARDED-FOR"))
		handler.ServeHTTP(w, r)
	})
}

func main() {
	var err error
	// Open our ConfigFile
	for _, v := range ConfigFileName {
		ConfigFile, err = os.Open(v)
		// if we os.Open returns an error then handle it
		if err != nil {
			log.Println(err)
			continue
		}
		defer ConfigFile.Close()
		log.Println("Successfully opened configuration file", v)
		confByteValue, _ := ioutil.ReadAll(ConfigFile)
		err := json.Unmarshal(confByteValue, &CFG)
		if err != nil {
			log.Println("JSONUnmarshal error:", err)
		}
		log.Println("Configuration read from file:", CFG)
		break
	}
	// Open our jsonFile
	JSONFile, err = os.OpenFile(CFG[0].DataJSONFile, os.O_RDWR, 0755)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully Opened", CFG[0].DataJSONFile)
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
