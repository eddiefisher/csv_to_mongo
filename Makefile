.PHONY: build test run

PROJECT?=home.dev/toster/csv_to_mongo/src
APP?=csv_to_mongo

RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
LDFLAGS?=-ldflags "-w -s \
	-X ${PROJECT}/version.Release=${RELEASE} \
	-X ${PROJECT}/version.BuildTime=${BUILD_TIME} \
	-X ${PROJECT}/version.Commit=${COMMIT}"

# Build the project
build:
	cd ./src ; go build ${LDFLAGS} -o ../bin/${APP}

# Test the project
test:
	cd ./src ; go test -v -race ./...
