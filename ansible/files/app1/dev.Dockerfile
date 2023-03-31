FROM golang:1.18.2-alpine
RUN apk add git gcc g++
ENV CGO_ENABLED 1
ENV GOPATH /go
ENV CC gcc


ADD . /go/src/app1
RUN cd /go/src/app1 && go mod tidy
WORKDIR /go/src/app1/api

RUN go install github.com/cosmtrek/air@latest
ENTRYPOINT air run main.go
EXPOSE 8080

