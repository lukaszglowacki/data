.ONESHELL:

all:
	docker-compose up -d --no-deps --build

local:
	go build -o ./worker ./cmd/worker/main.go 
	go build -o ./service ./cmd/service/main.go 


