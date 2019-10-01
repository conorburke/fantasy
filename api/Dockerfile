FROM golang:latest

RUN apt-get update
RUN apt-get install vim -y
WORKDIR /go/src/api

COPY . .
RUN rm -f go.mod
RUN rm -f go.sum
RUN go mod init github.com/conorburke/fantasy/api
RUN go build

CMD ["./api"]