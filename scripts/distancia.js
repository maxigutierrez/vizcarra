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
const fetch = require('node-fetch');
var Request = require('tedious').Request;
var Connection = require('tedious').Connection;
var config = {
    server: '10.8.1.9', //update me
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
    console.log('Start')
    connection.on('connect', function(err) {
        var request = new Request(`select a.*,b.latitud latitudo,b.longitud longitudo,c.latitud latitudd,c.longitud longitudd 
        from franquicias_distancia a inner join 
        franquicias b on a.origen_id = b.id  inner join
        franquicias c on a.destino_id = c.id 
where isnull(b.latitud,'') >'' and  isnull(c.latitud,'') >'' and distancia is null order by a.id desc`, function(err, rowCount) {
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
            console.log(origen[i]);
            const result = await medir(origen[i])
            console.log('ok',result)
            await wait(10)
        }
       process.exit(1);
    },
    function(err) {});


function  wait(seconds){
    var waitTill = new Date(new Date().getTime() + seconds * 1000);
    while(waitTill > new Date()){}
}


async function distancia(origen) {


    const googleMapsClient = require('@google/maps').createClient({
        key: 'AIzaSyAdYxa9N9WsRCQ2NLOj7f5V-TZckugW3bU'
    });
    try {
        let myLatLng = { lat: parseFloat(origen.latitudo), lng: parseFloat(origen.longitudo) }
        let myLatLngd = { lat: parseFloat(origen.latitudd), lng: parseFloat(origen.longitudd) }
        var request = {
            origin: myLatLng,
            destination: myLatLngd
        };
        await googleMapsClient.directions(request, function(err, response) {
            if (response.json.status == 'OK') {
                executeStatement1(`update franquicias_distancia set distancia= ` + response.json.routes[0].legs[0].distance.value + ' where id =' + origen.id)
                console.log(`update franquicias_distancia set distancia= ` + response.json.routes[0].legs[0].distance.value + ' where id =' + origen.id);
            } else {
                console.log('Error:', err)
            }
        });

    } catch (e) {
        console.log(e.toString());
    }
}
async function  medir(origen){
  var latOrigen = origen.longitudo
  var lonOrigen = origen.latitudo
  var latDestino = origen.latitudd
  var lonDestino = origen.longitudd
  var API_KEY = "AIzaSyAdYxa9N9WsRCQ2NLOj7f5V-TZckugW3bU"

  var distanciaURL = "https://maps.googleapis.com/maps/api/distancematrix/json?units=imperial&origins=" + latOrigen + "," + lonOrigen + "&destinations=" + latDestino + "," + lonDestino + "&key=" + API_KEY

  const res = await fetch( distanciaURL, {
    method: 'GET',
    headers: {
        'Content-Type': 'application/json'
    }
  });
  const data = await res.json();
  if (data.status=='OK'){
        console.log(data.rows[0].elements[0].distance.value);
        executeStatement1(`update franquicias_distancia set distancia= ` + data.rows[0].elements[0].distance.value + ' where id =' + origen.id)
        console.log(`update franquicias_distancia set distancia= ` + data.rows[0].elements[0].distance.value + ' where id =' + origen.id);

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
