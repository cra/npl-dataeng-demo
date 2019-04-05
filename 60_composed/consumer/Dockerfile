FROM golang:latest
RUN mkdir /app
WORKDIR /app
RUN go get \
  github.com/bitly/go-nsq \
  github.com/prometheus/client_golang/prometheus \
  github.com/prometheus/client_golang/prometheus/promauto \
  github.com/prometheus/client_golang/prometheus/promhttp

ADD . /app/
RUN go build -o npl-consumer .
