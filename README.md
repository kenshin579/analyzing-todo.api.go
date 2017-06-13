[![circleci](https://img.shields.io/badge/build-circleci-blue.svg)](https://circleci.com/gh/rctyler/todo.api.go)
[![CircleCI](https://img.shields.io/circleci/project/github/rctyler/todo.api.go.svg)](https://circleci.com/gh/rctyler/todo.api.go)
[![Coverage Status](https://coveralls.io/repos/github/rctyler/todo.api.go/badge.svg?branch=master)](https://coveralls.io/github/rctyler/todo.api.go?branch=master)
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
## Deploy locally
In the project's root directory, run the following commands:
1. `docker build -t rctyler/todo.api.go .`
2. `cd deploy`
3. `docker-compose up`
## Deploy API from an artifact
Download artifacts (todo-api-go.tar.gz and docker-compose.yml) from [https://circleci.com/gh/rctyler/todo.api.go](https://circleci.com/gh/rctyler/todo.api.go) and run the following commands:
1. `docker load < todo-api-go.tar.gz`
2. `docker-compose up`