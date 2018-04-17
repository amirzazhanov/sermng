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