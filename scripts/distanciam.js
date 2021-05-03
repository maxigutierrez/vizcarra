const fetch = require('node-fetch');
const sql = require('mssql')

async function sindistancias ()  {

      // make sure that any items are correctly URL encoded in the connection string
      await sql.connect('mssql://sa:Credifin2019@10.8.1.9/credifin')
      const result = await sql.query`select  a.*,b.latitud latitudo,b.longitud longitudo,c.latitud latitudd,c.longitud longitudd 
      from franquicias_distancia a inner join 
      franquicias b on a.origen_id = b.id  inner join
      franquicias c on a.destino_id = c.id 
where isnull(b.latitud,'') >'' and  isnull(c.latitud,'') >'' and distancia is null order by a.id desc`
      n= result.recordset.length;

      for (var i = 0; i < n; i++) {
        console.log(result.recordset[i]);
        const res = await medir(result.recordset[i],sql)

       console.log('ok')
//        await wait(10)
      }
//      process.exit(1)
}
  sindistancias();
function  wait(seconds){
    var waitTill = new Date(new Date().getTime() + seconds * 1000);
    while(waitTill > new Date()){}
}
async function  medir(origen,sql){
  var latOrigen = origen.latitudo
  var lonOrigen = origen.longitudo
  var latDestino = origen.latitudd
  var lonDestino = origen.longitudd
  var API_KEY = "AIzaSyAdYxa9N9WsRCQ2NLOj7f5V-TZckugW3bU"

  var distanciaURL = "https://maps.googleapis.com/maps/api/distancematrix/json?units=metric&origins=" + latOrigen + "," + lonOrigen + "&destinations=" + latDestino + "," + lonDestino + "&key=" + API_KEY
  console.log(distanciaURL)
  const res = await fetch( distanciaURL, {
    method: 'GET',
    headers: {
        'Content-Type': 'application/json'
    }
  });
  const data = await res.json();
  if (data.status=='OK' && data,data.rows[0],data.rows[0].elements[0].status!='ZERO_RESULTS' ){

    console.log(`update franquicias_distancia set distancia= ` + data.rows[0].elements[0].distance.value + ' where id =' + origen.id);
   // await sql.connect('mssql://sa:Credifin2019@10.8.1.9/credifin')
    //const s = await sql.query`update franquicias_distancia set distancia= ` + data.rows[0].elements[0].distance.value + ' where id =' + origen.id+';';
    const request = new sql.Request()
    request.query(`update franquicias_distancia set distancia= ` + data.rows[0].elements[0].distance.value + ' where id =' + origen.id+';').then(result => {
      console.log('pasa',result.rowsAffected)
    })
  }
}
