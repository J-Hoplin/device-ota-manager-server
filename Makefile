COMPILE_OUTPUT ?= server
IMAGE_TAG ?= ota-server
CONTAINER_NAME ?= ota-servr-api

compile:
	go build -o $(COMPILE_OUTPUT) ./cmd/main.go

run:
	./$(COMPILE_OUTPUT)

docker-build:
	docker build -t $(IMAGE_TAG) .

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-bash:
	docker exec -it $(CONTAINER_NAME) /bin/bash