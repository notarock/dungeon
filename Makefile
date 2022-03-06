##
# Dungeon
#
# @file
# @version 0.1

all: build test vet fmt lint build

run:
	go run main.go

build:
	go build -o bin/dungeon

test:
	go test ./...

vet:
	go vet ./...

fmt:
	gofmt -w .

# end
