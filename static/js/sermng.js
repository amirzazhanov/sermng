    function CreateTableFromJSON() {
        var myJson;
        jQuery.ajax( {
            url: 'http://localhost:8080/v1/records',
            type: 'GET',
            contentType: 'application/json; charset=utf-8',
            success: function( response ) {
                myJson = response;
                // EXTRACT VALUE FOR HTML HEADER. 
                // ('Book ID', 'Book Name', 'Category' and 'Price')
                var col = [];
                for (var i = 0; i < myJson.length; i++) {
                    for (var key in myJson[i]) {
                        if (col.indexOf(key) === -1) {
                            col.push(key);
                        }
                    }
                }
                // CREATE DYNAMIC TABLE.
                var table = document.createElement("table");
                table.setAttribute('class', 'table table-hover table-condensed');
                // ADD JSON DATA TO THE TABLE AS ROWS.
                for (var i = 0; i < myJson.length; i++) {
                    tr = table.insertRow(-1);
                    for (var j = 0; j < col.length; j++) {
                        var tabCell = tr.insertCell(-1);
                        tabCell.innerHTML = myJson[i][col[j]];
                    }
                }
            
                // FINALLY ADD THE NEWLY CREATED TABLE WITH JSON DATA TO A CONTAINER.
                var divContainer = document.getElementById("records_content");
                divContainer.innerHTML = "";
                divContainer.appendChild(table);
                }
          } );
        }
