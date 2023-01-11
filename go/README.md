# README

## this is a simple implementation of kafka producer -> broker -> consumer in golang

### requirements
install golang

### create network
```docker network create mynetwork```

### to run zookeeper

```docker run --name zookeeper  -p 2181:2181 -d zookeeper --network mynetwork```

### to run kafka broker

``` docker run -p 9092:9092 --name kafka -d --network mynetwork confluentinc/cp-kafka ```

### to stop and remove the containers 
``` docker stop zookeeper kafka ``` and ``` docker rm zookeeper kafka ```

### run the producer and consumer
```go run main.go```
