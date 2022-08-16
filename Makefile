include .env

build:
	go build -o .bin/main cmd/main.go
run: build
	./.bin/main
dev:
	go run cmd/main.go
pub:
	go run cmd/pub.go

migrate_init:
	migrate create -ext sql -dir ./schema -seq init
migrate:
	migrate -path ./schema -database 'postgres://postgres:manjaro_root@0.0.0.0:5436/wbl0?sslmode=disable' up