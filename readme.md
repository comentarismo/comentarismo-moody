# Comentarismo Moody

## This is the Comentarismo Sentiment Analysis tool that get the "feeling" form any text given the language is supported by the engine.

# Tokenizers
* Tokenizers are responsible for finding the Stop Words and breaking the text into pieces that can then be classified
* Eg: Portuguese Stopwords: http://www.ranks.nl/stopwords/portugese

# Classifier
* Classifiers are a set of words and weight, it uses the afinn http://neuro.imm.dtu.dk/wiki/AFINN

# Convey does not work for this project, open question: https://github.com/smartystreets/goconvey/issues/176
* $ goconvey --timeout=0

# Endpoints

## Moody - Sentiment Analysis on any Youtube link
http://localhost:3003/moody?vid=https%3A%2F%2Fwww.youtube.com%2Fwatch%3Fv%3DgKI78VOh4Aw

* Optional parameter lang=pt for portuguese videos


# Install:
```
// --- instal redis using yum ---
$ rpm -Uvh http://download.fedoraproject.org/pub/epel/6/x86_64/epel-release-6-8.noarch.rpm
$ rpm -Uvh http://rpms.famillecollet.com/enterprise/remi-release-6.rpm

$ yum --enablerepo=remi,remi-test install redis

[root@g7 ~]# service redis start/stop/restart
Usage: /etc/init.d/redis {start|stop|status|restart|condrestart|try-restart}
```


# Run server
```
REDIS_HOST=localhost REDIS_PORT=6379 RETHINKDB_HOST=ns351151.ip-91-121-75.eu RETHINKDB_PORT=28015 nohup ./main &.
REDIS_HOST=localhost REDIS_PORT=6379 RETHINKDB_HOST=91.121.75.229 RETHINKDB_PORT=28015 nohup ./main &.
```

# Tail logs
```
tail -f /opt/gocode/src/comentarismo-moody/nohup.out
```

# Godep
* $ godep go run main.go 
```
$ cat ~/.ssh/id_rsa.pub | ssh g7@g7-box "mkdir -m 755 -p ~/.ssh && cat >>  ~/.ssh/authorized_keys && chmod 644 ~/.ssh/authorized_keys"
```

I like the way info is displayed, showing as well common words and it detects emoticons from the text
I want that.


https://www.npmjs.com/package/sentimentjs
