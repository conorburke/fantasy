FROM golang:latest

RUN apt-get update
RUN apt-get install cron -y
RUN apt-get install vim -y
WORKDIR /go/src/etl

COPY . .
RUN rm -f go.mod
RUN rm -f go.sum
RUN go mod init github.com/conorburke/fantasy/etl
RUN go build

# RUN (crontab -l ; echo "0 * * * * /go/src/etl/etl") | crontab

CMD ["./etl"]
