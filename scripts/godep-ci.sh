#!/usr/bin/env bash

export GOPATH=/var/lib/jenkins/jobs/comentarismo-moody/workspace;
export PATH=$PATH:/var/lib/jenkins/jobs/comentarismo-moody/workspace/bin;

go get github.com/tools/godep;
go get github.com/stretchr/testify;
go get github.com/smartystreets/goconvey;
go get github.com/drewolson/testflight;
go get github.com/tsenart/vegeta;
cd src/comentarismo-moody;
godep restore;