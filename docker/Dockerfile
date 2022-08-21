FROM golang:1.19.0

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /go/src/go-starter
ENTRYPOINT [ "docker/entrypoint.sh" ]
