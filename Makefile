include .env

dev:
	go run main.go

build:
	CGO_ENABLED=0 go build -o ./bin/tago 
	upx ./bin/tago