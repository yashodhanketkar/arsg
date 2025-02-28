# INFO: Project variables
PROJECT_NAME := $(shell basename $(CURDIR))
.PHONY: run start build commit add test getcover cover clean

default: help

# INFO: Project commands
run:
	@go run main/main.go

start:
	@./build/${PROJECT_NAME}

startclean: build
	@./build/${PROJECT_NAME}

build: test clean 
	@go build -o ./build/${PROJECT_NAME} ./main/main.go

commit: add
	@git commit

add: test
	@git add .

test:
	@go test ./...

clean:
	@rm -rf ./build/

getcover: cover
	@go tool cover -html=./out/coverage.out -o ./out/coverage.html
	@powershell -Command "Start-Process '.\out\coverage.html'"

cover:
	@if not exist out mkdir out
	@go test -coverprofile=./out/coverage.out ./...

help:
	@echo "run       run the app"
	@echo "start     build and run the app"
	@echo "build     build the app"
	@echo "commit    commit the changes"
	@echo "add       add the changes"
	@echo "test      run tests"
	@echo "clean     clean the build directory"
	@echo "getcover  get the coverage report"
	@echo "cover     run the tests and generate the coverage report"
