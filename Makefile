##app commands
tidy:
	go mod download
	go mod tidy

run:
	go run main.go

## Docker environment commands
docker-start:
		docker-compose -f docker/docker-compose.yml up --build -d

docker-stop:
		docker-compose -f docker/docker-compose.yml down

docker-clear:
		docker stop $$(docker ps -a -q); docker container prune ; docker image prune ; docker volume prune;

#Publish test sns msg
publish-local-sns-test:
	aws --endpoint-url=http://localhost:4566 sns publish --topic-arn arn:aws:sns:us-east-1:000000000000:sns_sender_topic --message  'for test'
