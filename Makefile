.ONESHELL:

all:
	./build/run.sh

local:
	go build -o ./worker ./cmd/worker/main.go 
	go build -o ./service ./cmd/service/main.go 


