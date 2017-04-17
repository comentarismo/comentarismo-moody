#This will run the app
# http://software.clapper.org/daemonize/daemonize.html
# https://wiki.jenkins-ci.org/display/JENKINS/Spawning+processes+from+build

#authdisabled=true;
#limitdisabled=true;
#rediscachedisabled=true;
#GITHUB_KEY=c96ecc15dddd3d58753e
#GITHUB_SECRET=c35656e12f6120e21dd403a2e7c9fde80defe8d1
#APP_NAME='comentarismo-moody'
#WORKDIR='src/comentarismo-moody';
#GO_PROCESS='main.go';

#!/usr/bin/env bash

export GOPATH=/var/lib/jenkins/jobs/$APP_NAME/workspace;
export PATH=$PATH:/var/lib/jenkins/jobs/$APP_NAME/workspace/bin;

cd ${GOPATH}/${WORKDIR}
godep restore;

for PROCESS in ${GO_PROCESS}
do
pidf=${GOPATH}/${WORKDIR}_${PROCESS}.pid
exec 221>${pidf}
flock --exclusive --nonblock 221 ||
{
	echo "${pidf} already exists. so we will restart it for you ... ";
	kill `pidof $APP_NAME`
}
echo $$>&221
echo "${pidf} is not running, I am going to start an instance!!!";

godep go build -o ${GOPATH}/${WORKDIR}/$APP_NAME ${GOPATH}/${WORKDIR}/${GO_PROCESS}

/usr/local/sbin/daemonize -E BUILD_ID=dontKillMe -c ${GOPATH}/${WORKDIR} -p ${pidf} ${GOPATH}/${WORKDIR}/${APP_NAME} &

done