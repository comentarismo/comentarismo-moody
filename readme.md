$ goconvey --timeout=0

http://localhost:3000/?vid=https%3A%2F%2Fwww.youtube.com%2Fwatch%3Fv%3DgKI78VOh4Aw

$ godep go run main.go 


// --- instal redis using yum ---
$ rpm -Uvh http://download.fedoraproject.org/pub/epel/6/x86_64/epel-release-6-8.noarch.rpm
$ rpm -Uvh http://rpms.famillecollet.com/enterprise/remi-release-6.rpm

$ yum --enablerepo=remi,remi-test install redis

[root@g7 ~]# service redis start/stop/restart
Usage: /etc/init.d/redis {start|stop|status|restart|condrestart|try-restart}

Run server
``
REDIS_HOST=localhost REDIS_PORT=6379 db=ns351151.ip-91-121-75.eu:28015 nohup ./main &.
```

Tail logs
```
tail -f /opt/gocode/src/comentarismo-moody/nohup.out
```