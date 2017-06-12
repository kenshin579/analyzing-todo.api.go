[![CircleCI](https://img.shields.io/circleci/project/github/rctyler/todo.api.go.svg)](https://circleci.com/gh/rctyler/todo.api.go)
# todo.api.go
A simple TODO API built with Golang and Gin (this is an introductory project to learn Go). It is backed by RedisDB, and built using Docker.
## Build
In the project's root directory, run the following commands:
1. `docker build -t rctyler/todo.api.go .`
2. `cd deploy`
3. `docker-compose up`
