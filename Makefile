##
# Dungeon
#
# @file
# @version 0.1

all: test vet fmt lint build

test:
	go test ./...

vet:
	go vet ./...

fmt:
	gofumpt -w **/**.go

lint:
	go list ./... | xargs -L1 golint -set_exit_status

build:
	go build -o bin/dungeon

run:
	go run main.go

# end
