FROM golang:1.9-stretch

COPY ./deamon /go/src/gitlab.com/prj-risa/ex
WORKDIR /go/src/gitlab.com/prj-risa/ex

RUN go get ./
RUN go build

EXPOSE 8080