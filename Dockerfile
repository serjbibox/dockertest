FROM golang:latest AS build_stage
RUN mkdir -p go/src/dockertest
WORKDIR /go/src/dockertest
COPY ./ ./
RUN go env -w GO111MODULE=auto
RUN go install .
# ENTRYPOINT /go/bin/dockertest
EXPOSE 8080