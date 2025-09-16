# Kafka

```bash
docker compose up
```

Starts Confluent Platform, the docker compose file came from [here](https://raw.githubusercontent.com/confluentinc/cp-all-in-one/7.9.0-post/cp-all-in-one/docker-compose.yml), but with the following alterations:

 - remove all containers but `cp-enterprise-control-center` and `cp-server`

You can view the control center ui at http://localhost:9021. 

It creates a `test-user` topic and a `test-user-v8` consumer group on start - so the below Bento config should work after booting up the server!

### Basic Bento Config

```yaml
input:
  kafka:
    addresses: [localhost:9092]
    topics: [test-user]
    consumer_group: test-user-v8

output:
  kafka:
    addresses: [localhost:9092]
    topic: test-user
```

```sh
docker exec broker kafka-topics --bootstrap-server broker:29092 --create --topic test
```