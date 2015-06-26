FROM golang
MAINTAINER Douézan-Grard Guillaume - Quorums

ADD . /go/src/github.com/quorumsco/contacts

WORKDIR /go/src/github.com/quorumsco/contacts

RUN \
  go get && \
  go build

EXPOSE 8080

ENTRYPOINT ["./contacts"]
