.PHONY: start
start:
	docker compose up --build --remove-orphans

.PHONY: test
test:
	echo "Running commands in the container's terminal"
	@$(call docker_exec, 'cd /app/internal/utils && go test -v')
	@$(call docker_exec, 'cd /app/config && go test -v')
	@$(call docker_exec, 'cd /app/internal/handler && go test -v')

.PHONY: stop
stop:	
	docker compose down --rmi all

.PHONY: fmt
fmt:
	cd ./cmd && go fmt
	cd ./config && go fmt
	cd ./internal/handler && go fmt
	cd ./internal/logger && go fmt
	cd ./internal/utils && go fmt
	
.PHONY: tidy
tidy:
	go mod tidy

define docker_exec
	# Executes the given command/argument in the container's terminal	
	echo $(1)
	docker exec `docker ps --format '{{.Names}}' | grep matrixserver` sh -c  $(1)
endef