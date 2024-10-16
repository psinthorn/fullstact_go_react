# Build multi stage docker container 

# 1. Build golang container from base image and copy all the porject to WORKDIR in container 
# Get Go base image


FROM golang:1.23-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api  

# Optional 
RUN chmod +x /app/brokerApp

# 2. Build tiny docker container image from builder 
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/brokerApp /app

CMD ["/app/brokerApp"]





