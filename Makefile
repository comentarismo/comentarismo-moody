#!/usr/bin/env bash

default:
	scripts/gofmt_validate.sh;
	scripts/gotest.sh;

gofmt:
	scripts/gofmt_perform.sh;

test:
	scripts/gotest.sh;

vendor:
	@echo "--> Installing build dependencies"
	@godep save

.PHONY: all test

