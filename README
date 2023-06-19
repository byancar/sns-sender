# SNS Sender

This application demonstrates sending messages via SNS (Simple Notification Service) using AWS SDK for Go v2.

## App Commands

### Tidy

Run the following command to download and tidy Go modules:

```shell
go mod download
go mod tidy
```

### Run

Execute the following command to run the application:

```shell
go run main.go
```

## Docker Environment Commands

### Docker Start

Start the Docker environment by running the following command:

```shell
docker-compose -f docker/docker-compose.yml up --build -d
```

### Docker Stop

Stop the Docker environment by running the following command:

```shell
docker-compose -f docker/docker-compose.yml down
```

### Docker Clear

Stop all Docker containers, prune containers, images, and volumes using the following command:

```shell
docker stop $$(docker ps -a -q); docker container prune ; docker image prune ; docker volume prune;
```

## Publish Test SNS Message

To publish a test SNS message to a local endpoint, use the following command:

```shell
aws --endpoint-url=http://localhost:4566 sns publish --topic-arn arn:aws:sns:us-east-1:000000000000:sns_sender_topic --message 'for test'
```

Make sure to replace `arn:aws:sns:us-east-1:000000000000:sns_sender_topic` with the actual ARN of the SNS topic you want to use for testing.

