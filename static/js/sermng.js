var CONFIG = {
    "apiProtocol": "http",
    "apiPort": 8080,
    "apiHost": "localhost",
    "apiKey": "not_implemented"
    };
function displayRecords() {
    var myJson;
    jQuery.ajax( {
//        url: 'http://localhost:8080/v1/records',
        url: CONFIG.apiProtocol + '://' + CONFIG.apiHost + ':' + CONFIG.apiPort + '/v1/records',
        type: 'GET',
        contentType: 'application/json; charset=utf-8',
        success: function( response ) {
            myJson = response;
            // CREATE SCRIPT
            var script = document.createElement("script");
            script.setAttribute('type', 'text/javascript');
            script.innerHTML = '			function openallurls() {\n';
            // CREATE DYNAMIC TABLE.
            var table = document.createElement("table");
            table.setAttribute('class', 'table table-hover table-condensed');
            // ADD JSON DATA TO THE TABLE AS ROWS.
            for (var i = 0; i < myJson.length; i++) {
                tr = table.insertRow(-1);
                var tabCellDescr = tr.insertCell(-1);
                tabCellDescr.innerHTML = myJson[i]['description'];
                var tabCellCounter = tr.insertCell(-1);
                tabCellCounter.setAttribute('align', 'right');
                tabCellCounter.innerHTML =  '<div class="btn-group" role="group" aria-label="Count manager">'+
                                            '<button type="button" class="btn btn-secondary" onclick="editRecordCounter(' + myJson[i]['id'] + ', ' + ( myJson[i]['counter'] - 1 ) + ')" title="-1"><span class="oi oi-minus"></span></button>' +
                                            '<button type="button" class="btn btn-primary counter">' + myJson[i]['counter'] + '</button>' +
                                            '<button type="button" class="btn btn-secondary" onclick="editRecordCounter(' + myJson[i]['id'] + ', ' + ( myJson[i]['counter'] + 1 ) + ')" title="+1"><span class="oi oi-plus"></span></button>' +
                                            '</div>';
                var tabCellAct = tr.insertCell(-1);
                tabCellAct.setAttribute('align', 'right');
                tabCellAct.innerHTML =      '<div class="btn-group" role="group" aria-label="actions">' +
                                            '<a href="' + myJson[i]['url'] + '" class="btn btn-primary"><span class="oi oi-external-link"></span></a>'+
                                            '<button type="button" class="btn btn-primary"  data-toggle="modal" data-target="#update_record_modal" onclick="startEditRecord(' + myJson[i]['id'] + ')" title="Edit"><span class="oi oi-pencil"></span></button>' +
                                            '<button type="button" class="btn btn-primary" onclick="deleteRecord(' + myJson[i]['id'] + ')" title="Delete"><span class="oi oi-trash"></span></button>' +
                                            '</div>';
                script.innerHTML = script.innerHTML + 'window.open("' + myJson[i]['url'] + '", "' + myJson[i]['url'] + '");\n';
            }
            script.innerHTML = script.innerHTML + '}\n';
            // FINALLY ADD THE NEWLY CREATED TABLE WITH JSON DATA TO A CONTAINER.
            var divContainer = document.getElementById("records_content");
            divContainer.innerHTML = "";
            divContainer.appendChild(script);
            divContainer.appendChild(table);
            }
    } );
}
function deleteRecord(id) {
    jQuery.ajax( {
        url: CONFIG.apiProtocol + '://' + CONFIG.apiHost + ':' + CONFIG.apiPort + '/v1/records/' + id,
        type: 'DELETE',
        complete: function(result) {
            displayRecords();
        }
    });
}
function startEditRecord(id) {
    jQuery.ajax( {
        url: CONFIG.apiProtocol + '://' + CONFIG.apiHost + ':' + CONFIG.apiPort + '/v1/records/' + id,
        type: 'GET',
        contentType: 'application/json; charset=utf-8',
        success: function(response) {
            myJson = response;
            document.getElementById('update_record_description').value = myJson['description'];
            document.getElementById('update_record_counter').value = myJson['counter'];
            document.getElementById('update_record_url').value = myJson['url'];
            document.getElementById('hidden_record_id').value = id;
        },
        failure: function(errMsg) {
            alert(errMsg);
        }
    });
}
function updateRecordDetails() {
    jQuery.ajax({
        type: 'PUT',
        url: CONFIG.apiProtocol + '://' + CONFIG.apiHost + ':' + CONFIG.apiPort + '/v1/records/' + $('#hidden_record_id').val(),
        // The key needs to match your method's input parameter (case-sensitive).
        data: JSON.stringify({
            'description': $('#update_record_description').val(),
            'counter': parseInt($('#update_record_counter').val(), 10),
            'url': $('#update_record_url').val()
        }),
        contentType: 'application/json; charset=utf-8',
        dataType: 'json',
        complete: function(result) {
            displayRecords();
        },
        failure: function(errMsg) {
            alert(errMsg);
        }
    });    
}
function addRecord() {
    jQuery.ajax({
        type: 'POST',
        url: CONFIG.apiProtocol + '://' + CONFIG.apiHost + ':' + CONFIG.apiPort + '/v1/records',
        // The key needs to match your method's input parameter (case-sensitive).
        data: JSON.stringify({
            'description': $('#record_description').val(),
            'counter': parseInt($('#record_counter').val(), 10),
            'url': $('#record_url').val()
        }),
        contentType: 'application/json; charset=utf-8',
        dataType: 'json',
        complete: function(result) {
            displayRecords();
        },
        failure: function(errMsg) {
            alert(errMsg);
        }
    });    
}
function editRecordCounter(id, counterVal) {
    jQuery.ajax({
        type: "PUT",
        url: CONFIG.apiProtocol + '://' + CONFIG.apiHost + ':' + CONFIG.apiPort + '/v1/records/' + id,
        // The key needs to match your method's input parameter (case-sensitive).
        data: JSON.stringify({
            'counter': counterVal
        }),
        contentType: 'application/json; charset=utf-8',
        dataType: 'json',
        complete: function(result) {
            displayRecords();
        },
        failure: function(errMsg) {
            alert(errMsg);
        }
    });    
}