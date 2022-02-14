# syntax=docker/dockerfile:experimental

FROM golang:1 AS build-env
WORKDIR /
COPY . /growth-of-codes
ENV GO111MODULE on
RUN go install github.com/boyter/scc@latest
RUN cd /growth-of-codes && make bin/goc-analyze

FROM debian:stable-slim
LABEL org.opencontainers.image.source=https://github.com/STRRL/growth-of-codes
RUN apt-get update \
 && apt-get install -y --no-install-recommends ca-certificates
RUN update-ca-certificates

COPY --from=build-env /go/bin/scc /usr/local/bin/scc
COPY --from=build-env /growth-of-codes/bin/goc-analyze /usr/local/bin/goc-analyze
ENTRYPOINT ["goc-analyze"]
