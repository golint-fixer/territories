FROM golang
MAINTAINER Douézan-Grard Guillaume - Quorums

RUN go get github.com/quorumsco/contacts

ADD . /go/src/github.com/quorumsco/contacts

WORKDIR /go/src/github.com/quorumsco/contacts

RUN \
  go get -u && \
  go build

EXPOSE 8080

ENTRYPOINT ["./contacts"]
