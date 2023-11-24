build:
	@go build -o bin/activity

run: build
	./bin/activity

test:
	go test -v ./..