# syntax=docker/dockerfile:1
FROM golang:1.17.6-alpine3.15 as builder

MAINTAINER Peter Artemiev
ENV LANG en_US.UTF-8
ENV TZ UTC

RUN set -eux && \
    apk update && \
    apk upgrade && \
    apk add \
        bash \
        build-base \
        ca-certificates \
        curl \
        git \
        linux-headers

WORKDIR /src

COPY . /src

RUN go build \
    -a \
    -trimpath \
    -ldflags "-s -w" \
    -o app \
    "./L0.go"


FROM alpine:3.15.0

MAINTAINER Peter Artemiev
ENV LANG en_US.UTF-8
ENV TZ UTC

RUN apk upgrade --no-cache && \
    apk add --no-cache \
        ca-certificates && \
    mkdir -p /app/data

WORKDIR /app

COPY --from=builder /src/app .
COPY start.html .

CMD ["./app"]