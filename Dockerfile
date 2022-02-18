FROM golang:1.16.9-alpine3.14

ENV GOPATH=/go GO111MODULE=on CGO_ENABLED=1

RUN apk update && \
    apk --no-cache add git

COPY . /go/app
WORKDIR /go/app/

RUN go mod tidy && \
    go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]