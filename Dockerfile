FROM golang

ADD . /go/src/github.com/Quorumsco/sample-back

RUN \
  go install github.com/Quorumsco/sample-back

ENTRYPOINT /go/bin/sample-back migrate && /go/bin/sample-back serve

EXPOSE 8080
