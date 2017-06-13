[![CircleCI](https://img.shields.io/circleci/project/github/rctyler/todo.api.go.svg)](https://circleci.com/gh/rctyler/todo.api.go)
# todo.api.go
A simple TODO API built with Golang and Gin (this is an introductory project to learn Go). It is backed by RedisDB, and built in Docker.
## Install dependencies
1. `go get ./...`
## Build
In the project's root directory, run the following commands:
1. `cd src`
2. `go build`
## Run tests
1. `go test`
## Deploy locally
In the project's root directory, run the following commands:
1. `docker build -t rctyler/todo.api.go .`
2. `cd deploy`
3. `docker-compose up`
## Deploy API from an artifact
Download artifacts (todo-api-go.tar.gz and docker-compose.yml) from [https://circleci.com/gh/rctyler/todo.api.go](https://circleci.com/gh/rctyler/todo.api.go) and run the following commands:
1. `docker load < todo-api-go.tar.gz`
2. `docker-compose up`