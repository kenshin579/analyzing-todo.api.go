FROM golang

ADD . /go/src/github.com/rctyler/todoapp

RUN go get ./...

EXPOSE 8080

WORKDIR /go/src/github.com/rctyler/todoapp/src

CMD ["go", "run", "main.go"]
