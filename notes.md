# Notes

> notes taken during the course

<!-- https://gitignore.io -->
<!-- https://github.com/github/gitignore -->

https://github.com/edenhill/librdkafka

```sh
docker exec -it gokafka bash
go mod init fc-gokafka
go run cmd/producer/main.go
docker-compose ps
docker exec -it kafka_kafka_1 bash
kafka-topics --bootstrap-server=localhost:9092 --create --topic=TESTE --partitions=3
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=TESTE
```

https://github.com/edenhill/librdkafka/blob/master/CONFIGURATION.md
