BINARY_NAME=celeritas-skeleton
export GO111MODULE=on

build:
	@go get ./...
	@go mod tidy -e
	@echo "Building Celeritas..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Celeritas built!"

run: build
	@echo "Starting Celeritas..."
	@./tmp/${BINARY_NAME} &
	@echo "Celeritas started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

start-compose:
	docker compose up -d

stop-compose:
	docker compose down

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping Celeritas..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Celeritas!"

restart: stop start

## cover: opens coverage in browser
cover:
	@go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out -o=coverage.html

## cover: opens coverage in browser
cover-integration:
	@go test -coverprofile=coverage.out ./... --tags integration && go tool cover -html=coverage.out -o=coverage.html

## cover: opens coverage in browser
cover-unit:
	@go test -coverprofile=coverage.out ./... --tags unit && go tool cover -html=coverage.out -o=coverage.html