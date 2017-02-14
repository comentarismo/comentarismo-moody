#!/usr/bin/env bash

default:
	scripts/gofmt_validate.sh;
	scripts/gotest.sh;

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

permission:
	chmod +x scripts/godep-ci.sh;
	chmod +x scripts/gofmt_perform.sh;
	chmod +x scripts/gofmt_validate.sh;
	chmod +x scripts/gotest.sh;
	chmod +x scripts/start.sh;
	chmod +x scripts/stop.sh;

vendor-save:
	@echo "--> Installing build dependencies"
	@godep save

.PHONY: all test

