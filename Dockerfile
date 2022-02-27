FROM golang:1.16.9-alpine3.14

ENV GOPATH=/go GO111MODULE=on CGO_ENABLED=1 APPDIR=$GOPATH/src/gorl

RUN apk update && \
    apk --no-cache add git curl gcc musl-dev

COPY . $APPDIR/
WORKDIR $APPDIR

RUN go build -mod=vendor -o gorl app/main.go

ENTRYPOINT ["gorl/main"]
