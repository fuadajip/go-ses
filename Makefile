# Build And Development
start:
	@go run app/main.go

run: 
	@docker-compose up -d

stop:
	@docker-compose down

docker:
	@docker build -t registry.gitlab.com/mywishes/golang/go-ses:latest .

vendor:
	@go mod vendor

engine:
	@go build -o engine app/main.go

.PHONY: start run stop docker vendor engine 