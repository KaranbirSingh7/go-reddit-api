.DEFAULT_GOAL := help

run: 
	go run main.go

build: 
	go build -o bin/app main.go 

help: 
	echo "use: make run or make build"