include .env

.EXPORT_ALL_VARIABLES:

dev:
	go run -race main.go

build:
	CGO_ENABLED=0

	go build -ldflags="-s -w" -o ./bin/tago 
	upx ./bin/tago

prod: build
	./bin/tago

update:
	go get -u
	go mod tidy

.PHONY: dev build prod update race