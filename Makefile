install_system_pkgs_ubuntu:
	apt-get update
	apt install nodejs
	apt install npm
	apt install prometheus

install_node_deps:
	npm install nsqjs

install_go_deps:
	go get github.com/bitly/go-nsq
	go get github.com/prometheus/client_golang/prometheus
	go get github.com/prometheus/client_golang/prometheus/promauto
	go get github.com/prometheus/client_golang/prometheus/promhttp
	go get github.com/Shopify/sarama

