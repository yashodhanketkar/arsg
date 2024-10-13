.PHONY: test

test:
	@go test ./...

run:
	@go run main/main.go

commit: add
	@git commit

add: test
	@git add .

getcover: cover
	@go tool cover -html=./out/coverage.out -o ./out/coverage.html
	@powershell -Command "Start-Process '.\out\coverage.html'"

cover:
	@if not exist out mkdir out
	@go test -coverprofile=./out/coverage.out ./...
