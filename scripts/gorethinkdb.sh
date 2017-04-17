#!/usr/bin/env bash

table=comentarismo godep go test -v $(go list ./rethinkdb | grep -v /vendor/);
table=test godep go test -v $(go list ./rethinkdb | grep -v /vendor/);
