var http = require("http");

var handleRequest = function(request, response) {
  console.log('Received request for URL: ' + request.url)
  response.writeHead(200);
  response.end(`
<html>
  <head>
    <style>
      #message {
        text-align: center;
        margin-top: 100px;
        font-size: 22pt;
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
