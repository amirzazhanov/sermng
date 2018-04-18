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

* GET http://example.com/v1/records - get records from table (for now: ALWAYS GET ALL RECORDS)
* POST http://example.com/v1/records - record creation 
* PUT http://example.com/v1/records/{record_id} - record update (for now: counter update)
* DELETE http://example.com/v1/records/{record_id} - delete record

##TODO
1.[ ] full support of edit record (HTML/JS FE)
2. [ ] add configuration options for server (command line and config file)
  * [ ] HTML/JS FE directory root
  * [ ] REST API port
  * [ ] JSON file location
2. [ ] add JWT authentication to REST API (GO/REST BE + HTML/JS FE)
3. [ ] add support for multiuser (GO/REST BE + HTML/JS FE)