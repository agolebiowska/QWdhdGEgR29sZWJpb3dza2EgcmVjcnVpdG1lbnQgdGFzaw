DOCKER_COMPOSE_FILE		:= "docker-compose.yml"
PROJECT_CONTAINER_NAME	:= "weather_microservice"

# Starts `docker-compose`
.PHONY: build-start-containers
build-start-containers:
	docker-compose -f ${DOCKER_COMPOSE_FILE} up -d --build

## Run tests
#.PHONY: build-exec-containers
#build-exec-containers:
#	docker-compose -f ${DOCKER_COMPOSE_FILE} exec ${PROJECT_CONTAINER_NAME} go test ./...

# Removes all containes and all volumes
.PHONY: build-rm-containers
build-rm-containers:
	docker-compose -f ${DOCKER_COMPOSE_FILE} down -v
	docker-compose -f ${DOCKER_COMPOSE_FILE} rm -v

# Automates all the things
.PHONY: start
start: build-start-containers
#build-exec-containers

# Stops all the things
.PHONY: stop
stop:
	docker-compose -f ${DOCKER_COMPOSE_FILE} stop

# Clean everything
.PHONY: clean
clean: stop build-rm-containers