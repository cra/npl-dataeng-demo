const http = require("http");
const nsq = require('nsqjs');
const w = new nsq.Writer('127.0.0.1', 32809);
w.connect()

var handleRequest = function(request, response) {

  console.log('Received request for URL: ' + request.url)
  w.publish("my-topic", {dt: Date.now(), ua: request.headers['user-agent']});
  response.writeHead(200);
  response.end(`
<html>
  <head>
    <style>
      #message {
        text-align: center;
        margin-top: 100px;
        font-size: 18pt;
      }
    </style>
  </head>
  <body>
    <P id="message">Hello?</p>
  </body>
</html>`);
};


var www = http.createServer(handleRequest);
www.listen(8080);
