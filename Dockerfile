FROM golang:latest
ENV GOPATH /go
ENV GO111MODULE on

RUN GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.25.0

ADD Makefile /go/src/github.com/ltacker/ethboards-statechannels/Makefile
ADD go.mod /go/src/github.com/ltacker/ethboards-statechannels/go.mod
ADD pkg /go/src/github.com/ltacker/ethboards-statechannels/pkg
ADD cmd /go/src/github.com/ltacker/ethboards-statechannels/cmd
ADD .golangci.yml /go/src/github.com/ltacker/ethboards-statechannels/.golangci.yml

WORKDIR /go/src/github.com/ltacker/ethboards-statechannels
RUN make

EXPOSE 8000

ENTRYPOINT [ "ethboards-sc" ]