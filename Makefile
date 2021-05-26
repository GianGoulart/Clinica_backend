.PHONY: run watch mock docs test

VERSION = $(shell git branch --show-current)

# comandos para execução

run:
	VERSION=$(VERSION) go run main.go

run-watch:
	VERSION=$(VERSION) nodemon --exec go run main.go --signal SIGTERM

# comandos para teste

test:
	go test -coverprofile=coverage.out ./...

mock: 
	rm -rf ./mocks

	mockgen -source=./store/health/health.go -destination=./mocks/health_mock.go -package=mocks -mock_names=Store=MockHealthStore
	mockgen -source=./util/cache/cache.go -destination=./mocks/cache_mock.go -package=mocks

# comandos para documentação

docs:
	swag init