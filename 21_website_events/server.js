const http = require("http");
const nsq = require('nsqjs');
const w = new nsq.Writer('127.0.0.1', 32790);
w.connect()

const topic = "my-topic-site";

var handleRequest = function(request, response) {

  console.log('Received request for URL: ' + request.url)
  w.publish(topic, {
    dt: Date.now(),
    ua: request.headers['user-agent']
  });
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
    <P id="message">Hello?</p>
    <p>Этот сервер уже что-то пишет в nsq</p>
  </body>
</html>`);
};


var www = http.createServer(handleRequest);
www.listen(8080);
