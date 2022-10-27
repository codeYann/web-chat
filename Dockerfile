# syntax=docker/dockerfile:1

# It defines a golang version
FROM golang:1.19.1-alpine 

# Creating work directory
RUN mkdir /app

# Moving all files to WORKDIR folder
COPY . /app

# Set WORKDIR
WORKDIR /app

# Running go mod download to get all dependecies
RUN go build -o /main

CMD [ "/main" ]

EXPOSE 3333
