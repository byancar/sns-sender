#!/bin/bash

echo "installing jq"
apk update && apk add --no-cache jq

echo "configure aws"
aws configure set region "us-east-1"
aws configure set aws_access_key_id "FAKE"
aws configure set aws_secret_access_key "FAKE"

echo "setting up sns sqs"

################# CREATE LOCAL TOPIC AND QUEUE #################
TOPIC_NAME="sns_sender_topic"
QUEUE_NAME="settings"
TOPIC_ARN=$(aws --endpoint-url http://localhost:4566 sns create-topic --output text --name "$TOPIC_NAME")
QUEUE_URL=$(aws --endpoint-url http://localhost:4566 sqs create-queue --queue-name "$QUEUE_NAME" --output text)
QUEUE_ARN=$(aws --endpoint-url http://localhost:4566 sqs get-queue-attributes --queue-url "$QUEUE_URL" | jq -r ".Attributes.QueueArn")

aws --endpoint-url http://localhost:4566 sns subscribe --topic-arn "$TOPIC_ARN" --protocol sqs --notification-endpoint "$QUEUE_ARN" --output text

aws --endpoint-url http://localhost:4566 sns list-subscriptions
curl http://localhost:4566/health