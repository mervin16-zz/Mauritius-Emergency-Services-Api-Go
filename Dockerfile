ARG GO_VERSION=1.11

############################
# STEP 1 build executable binary
############################
FROM golang:${GO_VERSION}-alpine AS builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /mes
WORKDIR /mes

COPY . .
RUN go get -d -v
RUN go build -o ./mes ./main.go

############################
# STEP 2 build a small image
############################
FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

# Install bash
RUN apk add --no-cache bash

RUN mkdir -p /mes
WORKDIR /mes
COPY --from=builder /mes .

EXPOSE 8080

# Point the entrypoint to /app
ENTRYPOINT ["./mes"]