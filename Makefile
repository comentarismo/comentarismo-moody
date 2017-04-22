#!/usr/bin/env bash

default: gofmt
	make start;

gofmt:
	scripts/gofmt_perform.sh;

start: stop
	godep go build -o comentarismo-moody main.go
	nohup ./comentarismo-moody &

stop:
	pkill comentarismo-moody | true

status:
	ps -ef |grep comentarismo-moody

log:
	tail -f ./nohup.out

start-ci:
	scripts/godep-ci.sh
	scripts/start.sh

stop-ci:
	scripts/stop.sh

test:
	scripts/gotest.sh;

rethinkdb:
	scripts/gorethinkdb.sh;

goget:
	scripts/goget.sh;

permission:
	chmod +x scripts/godep-ci.sh;
	chmod +x scripts/gofmt_perform.sh;
	chmod +x scripts/gofmt_validate.sh;
	chmod +x scripts/gotest.sh;
	chmod +x scripts/start.sh;
	chmod +x scripts/stop.sh;
	chmod +x scripts/goget.sh

vendor-save:
	@echo "--> Installing build dependencies"
	@godep save

.PHONY: all test permission vendor-save

