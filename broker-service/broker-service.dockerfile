# Build multi stage docker container 

# 1. Build golang container from base image and copy all the porject to WORKDIR in container 
# Get Go base image

FROM golang:1.23-alpine as builder

RUN mkdir /app

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 go build -o brokerAPI ./cmd/api  

# Optional 
RUN chmod +x /app/brokerAPI

# 2. Build tiny docker container image from builder 
FROM alpine:latest

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/brokerAPI /app

CMD ["/app/brokerAPI"]





