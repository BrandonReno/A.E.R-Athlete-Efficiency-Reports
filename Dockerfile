FROM golang:1.16.4-buster
WORKDIR /
COPY . /go/src
RUN go get -d github.com/gorilla/mux
RUN go get -d github.com/go-playground/validator
CMD ["go", "run", "src/main.go"]