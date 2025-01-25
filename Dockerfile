FROM golang:1.23.5

RUN go version
ENV GOPATH=/

COPY ./ ./


RUN apt-get update
RUN apt-get -y install postgresql-client

