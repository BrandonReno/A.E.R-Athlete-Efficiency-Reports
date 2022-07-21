FROM golang:1.17.2-alpine as builder


RUN apk add --no-cache make git tzdata openssh

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download -x

COPY Makefile ./
COPY ./handler ./api
COPY ./models ./models 
COPY ./repository ./repository
COPY ./server ./server
COPY ./config ./config

RUN make build service=api output=server



FROM alpine:3.12
    RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/bin/api/server.bin ./

RUN  chmod +x server.bin
CMD /app/server.bin