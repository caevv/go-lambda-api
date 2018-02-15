FROM yunspace/serverless-golang:latest

RUN \
       go get github.com/DATA-DOG/godog/cmd/godog

