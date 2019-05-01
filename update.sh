#!/bin/sh

# if [ -d /tmp/mavlink ]; then
# 	echo Updating mavlink repo in /tmp
# 	(cd /tmp/mavlink; git pull)
# else
# 	echo Getting fresh clone of mavlink repo in /tmp
# 	(cd /tmp; git clone https://github.com/mavlink/mavlink.git)
# fi

#for dialect in /tmp/mavlink/message_definitions/v1.0/*.xml; do 
for dialect in ~/Project/mavlink/message_definitions/v1.0/*.xml; do 
	ddir=$(basename -s .xml $dialect | tr '[A-Z]' '[a-z]')
	go run mavgen/*.go $dialect > ${ddir}/mavlink.go
	(cd ${ddir}; go generate; go fmt; go build)
	echo
done