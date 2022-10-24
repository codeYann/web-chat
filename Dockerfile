# syntax=docker/dockerfile:1

# It defines a golang version
FROM golang:1.19.1-alpine 

RUN apk add --no-cache openssl

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

EXPOSE 3333

# Creating work directory
WORKDIR /app

# Creating web-chat folder and moving all files to it
RUN mkdir web-chat
RUN cd web-chat
COPY . web-chat

# Running go mod download to get all dependecies
RUN go mod tidy

ENTRYPOINT [ "go" "build" ]

RUN ./web-chat

