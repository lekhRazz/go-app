# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

RUN apk update && \
    apk add python3 g++ imagemagick make curl git

LABEL description="sample_go_application"
LABEL version="1.0"
LABEL Maintainer="Lekhrazz"


ARG WORKING_DIR=$HOME/sample_go_application

RUN mkdir -p $WORKING_DIR/
RUN mkdir -p $WORKING_DIR/build/

COPY go.mod $WORKING_DIR/
COPY go.sum $WORKING_DIR/

WORKDIR $WORKING_DIR/
RUN go mod tidy
COPY . $WORKING_DIR/
COPY .env.sample $WORKING_DIR/.env
RUN go build -o /build/docker_sample_go_application


CMD [ "/build/docker_sample_go_application" ]