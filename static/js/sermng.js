function displayRecords() {
    var myJson;
    jQuery.ajax( {
        url: 'http://localhost:8080/v1/records',
        type: 'GET',
        contentType: 'application/json; charset=utf-8',
        success: function( response ) {
            myJson = response;
            // CREATE DYNAMIC TABLE.
            var table = document.createElement("table");
            table.setAttribute('class', 'table table-hover table-condensed');
            // ADD JSON DATA TO THE TABLE AS ROWS.
            for (var i = 0; i < myJson.length; i++) {
                tr = table.insertRow(-1);
                var tabCellDescr = tr.insertCell(-1);
                tabCellDescr.innerHTML = myJson[i]['description'];
                var tabCellCounter = tr.insertCell(-1);
                tabCellCounter.innerHTML =  '<div class="btn-group" role="group" aria-label="Count manager">'+
                                            '<button type="button" class="btn btn-secondary"><span class="oi oi-minus"></span></button>' +
                                            '<button type="button" class="btn btn-primary counter">' + myJson[i]['counter'] + '</button>' +
                                            '<button type="button" class="btn btn-secondary"><span class="oi oi-plus"></span></button>' +
                                            '</div>';
                var tabCellAct = tr.insertCell(-1);
                tabCellAct.innerHTML =      '<div class="btn-group" role="group" aria-label="actions">' +
                                            '<a href="' + myJson[i]['url'] + '" class="btn btn-primary"><span class="oi oi-media-play"></span></a>'+
                                            '<button type="button" class="btn btn-primary" onclick="editRecord(' + myJson[i]['id'] + ')"><span class="oi oi-pencil"></span></button>' +
                                            '<button type="button" class="btn btn-primary" onclick="deleteRecord(' + myJson[i]['id'] + ')"><span class="oi oi-trash"></span></button>' +
                                            '</div>';
                                            
            }
        
            // FINALLY ADD THE NEWLY CREATED TABLE WITH JSON DATA TO A CONTAINER.
            var divContainer = document.getElementById("records_content");
            divContainer.innerHTML = "";
            divContainer.appendChild(table);
            }
    } );
}
function deleteRecord(id) {
    jQuery.ajax( {
        url: 'http://localhost:8080/v1/records/' + id,
        type: 'DELETE',
        success: function(result) {
            displayRecords();
        }
    });
}
function addRecord() {
    jQuery.ajax({
        type: "POST",
        url: "http://localhost:8080/v1/records",
        // The key needs to match your method's input parameter (case-sensitive).
        data: JSON.stringify({
            "description": $('#record_description').val(),
            "counter": parseInt($('#record_counter').val(), 10),
            "url": $('#record_url').val()
        }),
        contentType: "application/json; charset=utf-8",
        dataType: "json",
        success: function(result) {
            displayRecords();
        },
        failure: function(errMsg) {
            alert(errMsg);
        }
    });    
}