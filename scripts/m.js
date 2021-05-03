const {Client} = require("@googlemaps/google-maps-services-js");

const client = new Client({key: 'AIzaSyAdYxa9N9WsRCQ2NLOj7f5V-TZckugW3bU'});

client
  .geocode({
    params: {
      address: '1600 Amphitheatre Parkway, Mountain View, CA'
    },
    timeout: 1000, // milliseconds
  })
  .then((r) => {
    console.log(r.data);
  })
  .catch((e) => {
    console.log(e.response.data);
  });
