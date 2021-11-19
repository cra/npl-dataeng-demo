FROM golang:latest
RUN mkdir /app
WORKDIR /app

RUN go mod init mosyag.in/npl-demo
RUN go get \
  github.com/nsqio/go-nsq \
  github.com/prometheus/client_golang/prometheus \
  github.com/prometheus/client_golang/prometheus/promauto \
  github.com/prometheus/client_golang/prometheus/promhttp

ADD . /app/
RUN go build -o npl-consumer .
