FROM golang
MAINTAINER DANIEL JB - Quorums

RUN go get github.com/tools/godep

ADD . /go/src/github.com/quorumsco/territory

WORKDIR /go/src/github.com/quorumsco/territory

RUN godep go build

EXPOSE 8080

ENTRYPOINT ["./territory"]
