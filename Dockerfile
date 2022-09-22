# syntax=docker/dockerfile:1

# It defines a golang version
FROM golang:1.19.1-alpine 

# Creating work directory
WORKDIR /web-chat

# Moving all files to WORKDIR
COPY . .

# Running go mod download to get all dependecies
RUN go mod tidy

RUN go build
RUN ./web-chat






