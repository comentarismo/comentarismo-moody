#!/bin/bash
#!/bin/env go
# given script runs continuously every 5m checking for downtime
# */5 * * * * /opt/gocode/src/comentarismo-moody/main


cd /opt/gocode/src/comentarismo-moody
for PROCESS in main
do
pidf=/tmp/comentarismo-moody_${PROCESS}.pid
exec 221>${pidf}
flock --exclusive --nonblock 221 ||
{
        echo "${pidf} already exists. If you are sure ${PROCESS} is not runnig then delete ${pidf}"
        continue
}
echo $$>&221
echo "$PROCESS is not running, I am going to start an instance!!!";
REDIS_HOST=localhost REDIS_PORT=6379 db=ns351151.ip-91-121-75.eu:28015 nohup ./$PROCESS 2>&1 > /tmp/comentarismo-moody_$PROCESS.log &.

done
