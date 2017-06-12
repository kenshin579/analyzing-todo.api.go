FROM golang

ADD . /go/src/github.com/rctyler/todo.api.go

RUN go get ./...

EXPOSE 8080

WORKDIR /go/src/github.com/rctyler/todo.api.go/src

CMD ["go", "run", "main.go"]
