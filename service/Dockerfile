# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.19.1-alpine3.15 AS build

WORKDIR /app

COPY . ./
RUN go mod download

RUN go build -o /backendApi

##
## Deploy
##
FROM alpine:3.15.0

WORKDIR /

COPY --from=build /backendApi /backendApi

ENTRYPOINT ["/backendApi"]