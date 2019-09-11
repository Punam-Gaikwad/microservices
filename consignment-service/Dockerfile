FROM golang:1.9.0 as builder

WORKDIR /Go_Lang/Workspace/src/github.com/Punam-Gaikwad/microservices/consignment-service

COPY . .

FROM alpine:latest

RUN apk --no-cache add ca-certificates
RUN apk add --no-cache bash

RUN mkdir /app
WORKDIR /app
COPY --from=builder /Go_Lang/Workspace/src/github.com/Punam-Gaikwad/microservices/consignment-service .
CMD ["./consignment-service"]