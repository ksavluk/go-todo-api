WRKDIR := $(PWD)
DOCKER_IMAGE_NAME = github.com/ksavluk/go-todo-api
DOCKER_IMAGE_VERSION ?= develop

default: build

build:
	@ echo "-> Building binary..."
	go build -o bin/todo cmd/todo/main.go cmd/todo/app.go
.PHONY: build

docker-image-build:
	@ echo "-> Building docker image..."
	docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_VERSION) $(WRKDIR)
.PHONY: docker-image-build

start-env:
	@ echo "-> Starting test environment..."
	docker-compose up --build -d
.PHONY: start-env

stop-env:
	@ echo "-> Stopping test environment..."
	docker-compose down
.PHONY: stop-env