const http = require("http");
const nsq = require('nsqjs');
const w = new nsq.Writer('127.0.0.1', 4150);
w.connect()

var handleRequest = function(request, response) {

  console.log('Received request for URL: ' + request.url)
  w.publish("my-topic-site", {dt: Date.now(), ua: request.headers['user-agent'], source: "dockerized"});
  response.writeHead(200);
  response.end(`
<!doctype html>
<html lang="en">
<html>
  <head>
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
    <P id="message">Hello?</p>
    <p>Этот сервер уже что-то пишет в nsq</p>
  </body>
</html>`);
};


var www = http.createServer(handleRequest);
www.listen(8080);
