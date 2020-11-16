TOPIC?=my-topic
LOG_DIR?=/tmp

run_lookup:
	nsqlookupd

run_daemon:
	nsqd --lookupd-tcp-address=127.0.0.1:4160 --data-path=/tmp/tmp.Kf3VgFJqBF

run_admin:
	nsqadmin --lookupd-http-address=127.0.0.1:4161

stream_to_file:
	nsq_to_file --topic=${TOPIC} --output-dir=${LOG_DIR} --lookupd-http-address=127.0.0.1:4161

pub:
	curl -d 'test message' "http://127.0.0.1:4151/pub?topic=${TOPIC}"

read:
	tail -F ${LOG_DIR}/${TOPIC}*.log
