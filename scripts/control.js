/*
var sql = require('mssql');
var config = {
    server: '192.168.0.32',
    database: 'Credifin',
    user: 'sa',
    password: 'pp1234',
    port: 1433
};
*/
var Request = require('tedious').Request;
var Connection = require('tedious').Connection;
var config = {
    server: '10.8.1.1', //update me
    authentication: {
        type: 'default',
        options: {
            userName: 'sa', //update me
            password: 'Credifin2019' //update me
        }
    },
    options: {
        // If you are on Microsoft Azure, you need encryption:
        encrypt: true,
        database: 'Credifin' //update me
    }
};
var q = require('q');
module.exports = findAll;

function findAll() {
    var result = [];
    var deferred = q.defer(); // create a promise
    var connection = new Connection(config);
    connection.on('connect', function(err) {
        var request = new Request(`select id,latitud, longitud,cp
        from clientes_domicilios
         where isnull(latitud,'') >'' and  isnull(latitud,'') >''and cp>'' `, function(err, rowCount) {
            deferred.resolve(result); // resolve promise
        });
        request.on('row', function(columns) {
            var row = {};
            columns.forEach(function(column) {
                row[column.metadata.colName] = column.value;
            });
            result.push(row);
        });
        connection.execSql(request);

    });
    return deferred.promise; // return the promise for future
}


findAll().then(async function(origen) {
        console.log('start ', origen.length)
        for (i = 0; i < origen.length; i++) {

            const result = await distancia(origen[i])
            console.log('ok')
            sleep(1)
        }

    },
    function(err) {});

findAll().then(async function(origen) {
        console.log('start ', origen.length)
        for (i = 0; i < origen.length; i++) {
            const result = await distancia(origen[i])
            const sleep = (waitTimeInMs) => new Promise(resolve => setTimeout(resolve, waitTimeInMs));
            await sleep(500);
        }
    },
    function(err) {});


function distancia(origen) {
    const googleMapsClient = require('@google/maps').createClient({
        key: 'AIzaSyDhMami5rfzvzC4-kYaVBSrpVXES1MSZDk'
    });
    // 'AIzaSyDQs3WrPDLZ7mssC033PJh2XRL-ZaGsouI'
    try {
        let myLatLng = { lat: parseFloat(origen.latitud), lng: parseFloat(origen.longitud) }
        console.log(myLatLng)
        var latlng = new googleMapsClient.LatLng(parseFloat(origen.latitud), parseFloat(origen.longitud));
        googleMapsClient.geocode({ 'latLng': latlng }, function(results, status) {
            if (status == google.maps.GeocoderStatus.OK) {
                if (results[1]) {
                    console.log(results)
                } else {
                    alert("No results found");
                }
            } else {
                alert("Geocoder failed due to: " + status);
            }
        })

    } catch (e) {
        console.log(e.toString());
    }
}

function executeStatement1(sql) {
    var connection = new Connection(config);
    connection.on('connect', function(err) {
        request = new Request(sql, function(err) {
            if (err) {
                console.log(err);
            }
        });
        connection.execSql(request);
    });
}