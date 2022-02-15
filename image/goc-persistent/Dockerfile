# syntax=docker/dockerfile:experimental

FROM golang:1 AS build-env
WORKDIR /
COPY . /growth-of-codes
ENV GO111MODULE on
RUN cd /growth-of-codes && make bin/goc-persistent

FROM debian:stable-slim
LABEL org.opencontainers.image.source=https://github.com/STRRL/growth-of-codes
RUN apt-get update \
 && apt-get install -y --no-install-recommends ca-certificates
RUN update-ca-certificates

COPY --from=build-env /growth-of-codes/bin/goc-persistent /usr/local/bin/goc-persistent
ENTRYPOINT ["goc-goc-persistent"]
