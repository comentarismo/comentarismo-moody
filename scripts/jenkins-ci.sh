#!/usr/bin/env bash

export GOPATH=/var/lib/jenkins/jobs/comentarismo/workspace;
export PATH=$PATH:/var/lib/jenkins/jobs/comentarismo/workspace/bin;
go get github.com/tools/godep;
go get github.com/stretchr/testify;
go get github.com/smartystreets/goconvey;
go get github.com/drewolson/testflight;
go get github.com/tsenart/vegeta;
cd src/comentarismo;
godep restore;

#This will run all tests
#authdisabled=true limitdisabled=true rediscachedisabled=true godep go test -v $(go list ./... | grep -v /vendor/);

godep go test -v $(go list ./rethinkdb | grep -v /vendor/);
authdisabled=true limitdisabled=true rediscachedisabled=true  godep go test -v $(go list ./server | grep -v /vendor/);
authdisabled=true rediscachedisabled=true godep go test -v $(go list ./limiter | grep -v /vendor/);
