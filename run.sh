#!/bin/bash

cd ./dockers/mysql/
docker-compose up -d

cd ../rabbitmq
docker-compose up -d 

cd ../../server
go build .
./server &

cd ../workers/
docker build -t pentest .
docker run -ti pentest /bin/bash
