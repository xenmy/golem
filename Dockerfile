# Build Container
FROM golang:latest as builder

WORKDIR /go/src/github.com/xenmy/golem
COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -o app main.go

# Runtime Container
FROM alpine:3.13.2
COPY --from=builder /go/src/github.com/xenmy/golem/app /app
COPY --from=builder /go/src/github.com/xenmy/golem/config.yaml /config.yaml

ENV SERVER_ADDRESS="0.0.0.0" \
    SERVER_PORT="8080"

ENTRYPOINT ["/app", "-e", "config"]
