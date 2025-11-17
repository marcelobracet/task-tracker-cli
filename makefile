run: 
	go run cmd/main.go

test:
	go test ./...

build:
	go build -o ./build/task-tracker cmd/main.go

run-build:
	./build/task-tracker