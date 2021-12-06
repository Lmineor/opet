FROM golang:1.16-alpine as builder

ENV GOPATH /go

COPY . $GOPATH/src/github.com/opet

WORKDIR $GOPATH/src/github.com/opet

RUN GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -ldflags "-s -w" -o server

FROM centos:7.5.1804

RUN mkdir -p /opet/bin

COPY --from=builder  /go/src/github.com/opet/server /opet/bin
COPY --from=builder  /go/src/github.com/opet/run.py /opet

RUN chmod +x /opet/run.py
ENTRYPOINT ["/opet/run.py"]
