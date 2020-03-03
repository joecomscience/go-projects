FROM golang:1.13.8-alpine as build
LABEL maintainer="joewalker"
ENV GOOS=linux
ENV GOARCH=amd64

COPY . .
RUN go mod download && \
    go build cmd/main.go

FROM alpine:3.11 as release

LABEL maintrainer="joewalker"
COPY --from=build /app/main /app
EXPOSE 8080
CMD ["./app"]