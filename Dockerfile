FROM golang:1.17.2-alpine as builder


RUN apk add --no-cache make git tzdata openssh

WORKDIR /app

COPY ./backend ./backend
RUN cd ./backend && make build service=. output=server



FROM alpine:3.12
    RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/backend/bin ./

RUN  chmod +x server.bin
CMD /app/server.bin