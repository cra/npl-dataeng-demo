COMPONENT=newprolab-service-app
VERSION?=$(shell md5sum ./server.js | cut -f 1 -d " ")

build:
	docker build -t ${COMPONENT}:${VERSION} .

run:
	docker run -it --net host ${COMPONENT}:${VERSION}
