var request = require('request');
var S = require('string');

var url = 'http://api.icndb.com/jokes/random';

request({ url: url, 'json': true} , function (error, response, body) {
  if (!error && response.statusCode == 200) {
      console.log(S(body.value.joke).decodeHTMLEntities().s);
  }
});
