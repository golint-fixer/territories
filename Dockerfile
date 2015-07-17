FROM golang
MAINTAINER Douézan-Grard Guillaume - Quorums

RUN go get github.com/tools/godep

ADD . /go/src/github.com/quorumsco/contacts

WORKDIR /go/src/github.com/quorumsco/contacts

RUN \
  godep restore && \
  go build

EXPOSE 8080

ENTRYPOINT ["./contacts"]
