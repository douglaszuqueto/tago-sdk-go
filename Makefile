include .env

dev:
	go run main.go

prod: build
	./bin/tago

build:
	CGO_ENABLED=0 go build -o ./bin/tago 
	upx ./bin/tago

race:
	go run -race main.go