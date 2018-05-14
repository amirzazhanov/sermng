// Package sermng to test server implementation
package main

// https://golang.org/doc/articles/wiki/
// http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/
// https://thenewstack.io/make-a-restful-json-api-go/

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

// JSONFile GLOBAL FILE pointer to main JSON FILE
var JSONFile *os.File

// ConfigFile Global File pounter to configuration file
var ConfigFile *os.File

// ConfigFileName List of default locations for configuration filenames
var ConfigFileName = []string{"sermngconfig.json", "/etc/sermngconfig.json", "/usr/local/etc/sermngconfig.json"}

// Conf describes configuration file format
type Conf struct {
	RestAPIPort        uint32 `json:"restAPIPort,omitempty"`
	RestAPIBindAddress string `json:"restAPIBindAddress,omitempty"`
	DataJSONFile       string `json:"dataJSONFile,omitempty"`
}

// CFG Configuration to parse
var CFG []Conf

// Log function
func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s\t%s\t%s\t%s\tX-FORWARDED-FOR: %s", r.RemoteAddr, r.Method, r.Proto, r.URL.Path[1:], r.Header.Get("X-FORWARDED-FOR"))
		handler.ServeHTTP(w, r)
	})
}

func main() {
	var err error
	var dir string

	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the dir 'static'")
	flag.Parse()
	log.Println("base directory for static html/js/css files:", dir)
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
		log.Println("Address string:", CFG[0].RestAPIBindAddress+":"+strconv.FormatUint(uint64(CFG[0].RestAPIPort), 10))
		log.Println("JSON data file location:", CFG[0].DataJSONFile)
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
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.FileServer(http.Dir(dir)))
	r.HandleFunc("/v1/records", CreateRecord).Methods("POST")
	r.HandleFunc("/v1/records/{record_id:[0-9]+}", ReadRecord).Methods("GET")
	r.HandleFunc("/v1/records", ReadAllRecords).Methods("GET")
	r.HandleFunc("/v1/records/{record_id:[0-9]+}", UpdateRecord).Methods("PUT")
	r.HandleFunc("/v1/records/{record_id:[0-9]+}", DeleteRecord).Methods("DELETE")
	log.Fatal(http.ListenAndServe(CFG[0].RestAPIBindAddress+":"+strconv.FormatUint(uint64(CFG[0].RestAPIPort), 10), Log(r)))
}
