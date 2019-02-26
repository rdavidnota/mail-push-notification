FROM golang:alpine

RUN apk update \
    && apk add --virtual build-dependencies \
        build-base \
        gcc \
        wget \
        git \
    && apk add \
        bash

RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/app
COPY [".", "/go/src/app"]

RUN dep ensure

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go", "run", "main.go"]