FROM golang:1.13.8-alpine as build
LABEL maintainer="joewalker"
ENV GOOS=linux
ENV GOARCH=amd64

COPY . .
RUN go mod download && \
    go build main.go
CMD ["./main"]