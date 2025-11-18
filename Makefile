.PHONY: all
all: compile

.PHONY: compile
compile:
	solc --abi --bin --base-path eth/node_modules --overwrite eth/skavenge.sol -o eth/build
	abigen --bin eth/build/Skavenge.bin --abi eth/build/Skavenge.abi --pkg bindings --type Skavenge --out eth/bindings/bindings.go

# Docker targets
.PHONY: docker-build
docker-build:
	docker compose build

.PHONY: docker-up
docker-up:
	docker compose up -d hardhat

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