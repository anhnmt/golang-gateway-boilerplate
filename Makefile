DOCKER_COMPOSE=./build/docker-compose.yml

docker.build:
	docker compose -f $(DOCKER_COMPOSE) build --push

docker.up:
	docker compose -f $(DOCKER_COMPOSE) up -d --force-recreate --pull always

docker.down:
	docker compose -f $(DOCKER_COMPOSE) down

wire.gen:
	wire ./...

go.install:
	go install github.com/google/wire/cmd/wire@v0.5.0