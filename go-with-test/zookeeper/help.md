# go-zk

# install zookeeper

```bash
wget https://downloads.apache.org/zookeeper/stable/apache-zookeeper-3.8.3-bin.tar.gz
tar -zxvf apache-zookeeper-3.8.3-bin.tar.gz
cd apache-zookeeper-3.8.3-bin/conf
./bin/zkServer.sh start
```

# zkCli

./bin/zkCli.sh

```bash
create /go
create /go/services
create /go/services/localhost:8081
```

[zk: localhost:2181(CONNECTED) 3] ls -R /go
/go
/go/services
/go/services/localhost:8081

[zk: localhost:2181(CONNECTED) 6] set /go "hello go"
[zk: localhost:2181(CONNECTED) 7] get /go
hello go
