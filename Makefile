.PHONY: all compile hardhat test

all: compile

compile:
	solc --abi --bin --base-path eth/node_modules --overwrite eth/skavenge.sol -o eth/build
	abigen --bin eth/build/Skavenge.bin --abi eth/build/Skavenge.abi --pkg bindings --type Skavenge --out eth/bindings/bindings.go

# Docker targets
.PHONY: docker-build
docker-build:
	docker compose build

.PHONY: rebuild-webapp
rebuild-webapp:
	docker compose build webapp
	@echo "Webapp image rebuilt"

.PHONY: rebuild-webapp-no-cache
rebuild-webapp-no-cache:
	docker compose build --no-cache webapp
	@echo "Webapp image rebuilt (no cache)"

.PHONY: docker-up
docker-up:
	docker compose up -d hardhat

.PHONY: start
start: start-with-setup

.PHONY: start-services
start-services:
	docker compose up -d hardhat webapp
	@echo "Services starting..."
	@echo "Hardhat: http://localhost:8545"
	@echo "Webapp: http://localhost:8080"

.PHONY: start-with-setup
start-with-setup:
	docker compose up -d hardhat
	@echo "Waiting for Hardhat to be ready..."
	@sleep 5
	docker compose up --abort-on-container-exit setup
	docker compose up -d webapp
	@echo "Services started with contract deployment..."
	@echo "Hardhat: http://localhost:8545"
	@echo "Webapp: http://localhost:8080"

.PHONY: stop
stop: stop-services

.PHONY: stop-services
stop-services:
	docker compose stop hardhat webapp

.PHONY: docker-test
docker-test:
	docker compose up --abort-on-container-exit test

.PHONY: docker-down
docker-down:
	docker compose down

.PHONY: docker-clean
docker-clean:
	docker compose down -v --rmi all

.PHONY: test-local
test-local: docker-build docker-up
	@echo "Waiting for Hardhat to be ready..."
	@sleep 5
	docker compose up --abort-on-container-exit test
	docker compose down

.PHONY: setup
setup:
	@echo "Running contract setup..."
	docker compose up --abort-on-container-exit setup

.PHONY: setup-local
setup-local:
	@echo "Running contract setup locally..."
	@if [ ! -f test-config.json ]; then \
		echo "Error: test-config.json not found"; \
		exit 1; \
	fi
	go run ./cmd/setup
