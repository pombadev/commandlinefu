clean:
	go clean

build:
	go build

sbuild:
	go build -ldflags "-s -w"

run:
	go run .