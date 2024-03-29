#!/usr/bin/env bash

default: gofmt
	make start;

gofmt:
	scripts/gofmt_perform.sh;

start: stop
	godep go build -o comentarismo-moody main.go
	RESTART_TIMEOUT_FACTOR=5000 RESTART_TIMEOUT_ENABLED=false nohup ./comentarismo-moody &
	make log

start-prod: stop
	 godep go build -o comentarismo-moody main.go
	RETHINKDB_HOST=localhost RETHINKDB_PASSWORD=test123 nohup ./comentarismo-moody &
	make log

stop:
	pidof comentarismo-moody |awk '{print $1}'| xargs kill | true;

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

rethinkdb-prod:
	RETHINKDB_HOST=localhost RETHINKDB_PASSWORD=  scripts/gorethinkdb.sh;

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

