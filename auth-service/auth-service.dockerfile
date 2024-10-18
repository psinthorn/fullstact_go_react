# Build multi stage docker container 

# 1. Build golang container from base image and copy all the porject to WORKDIR in container 
# Get Go base image

# FROM golang:1.23-alpine as builder

# RUN mkdir /app

# WORKDIR /app

# COPY . /app

# RUN CGO_ENABLED=0 go build -o authAPI ./cmd/api  

# # Optional 
# RUN chmod +x /app/authAPI

# 2. Build tiny docker container image from builder 
# *Remark: this containe use build image from Makefile
FROM alpine:latest

RUN mkdir /app

COPY authAPI /app

CMD ["/app/authAPI"]

