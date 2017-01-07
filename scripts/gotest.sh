#!/usr/bin/env bash

godep go test -v $(go list ./model | grep -v /vendor/);
godep go test -v $(go list ./lang | grep -v /vendor/);
godep go test -v $(go list ./sentiment | grep -v /vendor/);
godep go test -v $(go list ./server | grep -v /vendor/);
godep go test -v $(go list ./test | grep -v /vendor/);
