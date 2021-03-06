[![circleci](https://img.shields.io/badge/build-circleci-blue.svg)](https://circleci.com/gh/rctyler/todo.api.go)
[![CircleCI](https://img.shields.io/circleci/project/github/rctyler/todo.api.go/master.svg)](https://circleci.com/gh/rctyler/todo.api.go?branch=master)
[![Coveralls branch](https://img.shields.io/coveralls/rctyler/todo.api.go/master.svg)](https://coveralls.io/github/rctyler/todo.api.go)
# todo.api.go
A simple TODO API built with Golang and Gin (this is an introductory project to learn Go). It is backed by RedisDB, and built in Docker.
## Install dependencies
In the project's root directory, run the following commands:
1. `go get ./...`
## Build
In the project's root directory, run the following commands:
1. `cd src`
2. `go build`
## Run tests
In the project's root directory, run the following commands:
1. `cd tests`
2. `go test`
## Run code coverage
In the project's root directory, run the following commands:
1. `$GOPATH/bin/goveralls -v -service=circle-ci -repotoken 3QBLcfjGKHMtQevf0Q0lkADeRpUdscmdS -ignore src/main.go,src/data/redis.go,tests/mocks/mocks.go`

**2017-6-16 NOTE:** There is a bug somewhere in my CI environment where running code coverage on the build server is _not_ publishing to coveralls.io. However, running the command on the source code locally works.
## Deploy locally
In the project's root directory, run the following commands:
1. `docker build -t rctyler/todo.api.go .`
2. `cd deploy`
3. `docker-compose up`
## Deploy API from an artifact
Download artifacts (todo-api-go.tar.gz and docker-compose.yml) from [https://circleci.com/gh/rctyler/todo.api.go](https://circleci.com/gh/rctyler/todo.api.go) and run the following commands:
1. `docker load < todo-api-go.tar.gz`
2. `docker-compose up`
