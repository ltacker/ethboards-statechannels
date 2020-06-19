FROM golang:latest
ENV GOPATH /go
ENV GO111MODULE on

RUN GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.25.0

ADD go.mod /go/src/github.com/ltacker/ethboards-statechannels/go.mod
ADD Makefile /go/src/github.com/ltacker/ethboards-statechannels/Makefile
ADD pkg /go/src/github.com/ltacker/ethboards-statechannels/pkg
ADD cmd /go/src/github.com/ltacker/ethboards-statechannels/cmd

WORKDIR /go/src/github.com/ltacker/ethboards-statechannels/cmd/ethboards-sc
RUN go install

EXPOSE 8546

# ENV WEB3_HOST
# ENV WEB3_PORT
# ENV MONGO_HOST
# ENV MONGO_PORT
# ENV MONGO_USERNAME
# ENV MONGO_PASSWORD

ENTRYPOINT [ "ethboards-sc" ]