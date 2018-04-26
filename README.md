# sermng
1. manage single table JSON database with web app and REST API
2. do not use 3rd party http router (Just for fun)

## JSON FORMAT
id,description,counter,url

* id - numeric record ID (autoincrement)
* description - string description
* counter - numeric (should support increment/decrement operations)
* url - resource URL

## REST API

* GET http://example.com/v1/records/{record_id} - get records from table
* POST http://example.com/v1/records - record creation 
* PUT http://example.com/v1/records/{record_id} - record update (for now: counter update)
* DELETE http://example.com/v1/records/{record_id} - delete record

## TODO
- [x] full support of edit record (HTML/JS FE + GO/REST BE)
  - [x] add support for "GET single record" to REST API
  - [x] edit record support in HTML/JS FE
- [x] add configuration options for server and client (command line and config file)
  - [x] GO/REST BE (use JSON as config file format)
    - [x] \(config option) REST API port
    - [x] \(config option) REST API bind address
    - [x] \(config option) JSON file location
  - [x] HTML/JS FE
- [ ] add JWT authentication to REST API (GO/REST BE + HTML/JS FE)
  - [ ] learn github.com/dgrijalva/jwt-go
- [ ] add support for multiuser (GO/REST BE + HTML/JS FE)
