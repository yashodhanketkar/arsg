# INFO: Project variables
.PHONY: run start build commit add test getcover cover clean

default: help

# INFO: Project commands
run:
	@go run src/main.go ui --mode dev

run-rest:
	@go run src/main.go rest --mode dev

start:
	@./build/arsg

startclean: build
	@./build/arsg

build: test clean
	@mkdir -p build
	@go build -buildmode=exe -o ./build/arsg -trimpath src/main.go
	@upx --best --lzma ./build/arsg

build-dev: clean
	@mkdir -p build
	@go build -o build/arsg  src/main.go

install: build
	@chmod a+x ./scripts/install.sh
	@./scripts/install.sh

uninstall:
	@chmod a+x ./scripts/install.sh
	@./scripts/uninstall.sh

test:
	@go test ./...

cover:
	@mkdir -p out
	@go test -coverprofile=./out/coverage.out ./...


getcover: cover
	@go tool cover -html=./out/coverage.out -o ./out/coverage.html

clean:
	@rm -rf ./build/

help:
	@echo "run       run the app"
	@echo "run-rest  run the rest api"
	@echo "start     start the app"
	@echo "startclean start the app and clean the build folder"
	@echo "build     build the app"
	@echo "build-dev build the app for development"
	@echo "install   install the app"
	@echo "uninstall uninstall the app"
	@echo "test      run the tests"
	@echo "cover     run the tests and get the coverage report"
	@echo "getcover  get the coverage report"
	@echo "clean     clean the build folder"
	@echo "help      show this help message"
