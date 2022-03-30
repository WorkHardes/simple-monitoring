DOCKERFILE_PATH=.

DOCKER_IMAGE_NAME=simple-monitoring

help:
	@echo "Commands:"
	@echo "build: Compiles project"
	@echo "start-server: starts server on localhost and 8000 port"
	@echo "build-docker-image: builds docker image with name $(DOCKER_IMAGE_NAME)"
	@echo "run-docker-imange: runs image $(DOCKER_IMAGE_NAME)"
	@echo "lint: code linter with golangci"

lint:
	golangci-lint run

build:
	go build ./cmd/app/main.go

start-server:
	go run ./cmd/app/main.go

build-docker-image:
	docker build -t $(DOCKER_IMAGE_NAME) $(DOCKERFILE_PATH)

run-docker-image:
	docker run -it $(DOCKER_IMAGE_NAME)
