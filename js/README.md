# README

## this is a simple implementation of kafka producer -> broker -> consumer in javascript

### requirements
install node and run ```npm install``` in the dir js

### to run zookeeper

```docker run --name zookeeper  -p 2181:2181 -d zookeeper```

### to run kafka broker

``` docker run -p 9092:9092 --name kafka  -e KAFKA_ZOOKEEPER_CONNECT=hostname:2181 -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://hostname:9092 -e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 -d confluentinc/cp-kafka ```

### to stop and remove the containers 
``` docker stop zookeeper kafka ``` and ``` docker rm zookeeper kafka ```
