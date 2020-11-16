const http = require("http");
const nsq = require('nsqjs');
const w = new nsq.Writer('nsqd', 4150);
w.connect()

var handleRequest = function(request, response) {

  console.log('Received request for URL: ' + request.url)
  w.publish("my-topic", {
    dt: Date.now(),
    ua: request.headers['user-agent'],
    source: "composed"
    }
  );
  response.writeHead(200);
  response.end(`
<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <style>
      #message {
        text-align: center;
        margin-top: 100px;
        font-size: 18pt;
      }
      p {
        text-align: center;
        font-size: 12pt;
      }
    </style>
  </head>
  <body>
    <P id="message">Hello!</p>
    <p>Этот сервер прямиком из docker-compose.</p>
  </body>
</html>`);
};


var www = http.createServer(handleRequest);
www.listen(8080);
