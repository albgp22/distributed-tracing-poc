all: start

.PHONY: type-a-service
type-a-service:
	@docker build -t type-a-service -f type-a-service/Dockerfile .

.PHONY: type-b-service
type-b-service:
	@docker build -t type-b-service -f type-b-service/Dockerfile .

.PHONY: start
start: type-a-service type-b-service
	@docker compose up --remove-orphans

.PHONY: stop
stop:
	@docker rmi type-a-service type-b-service --force