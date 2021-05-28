FROM golang:alpine as Base

#Download git for go get to grab dependencies
RUN apk update && apk add --no-cache git

#Set the working directory in the container
WORKDIR GOPATH$/src/AERpkg/A.E.R/

#Copy everything from the local dir into the workdir
COPY . .

#Get all the dependencies
RUN go get -d -v

#Compile the server
RUN go build -o /go/bin/AERserver

FROM scratch
COPY --from=Base /go/bin/AERserver go/bin/AERserver

EXPOSE 9090

CMD [ "go/bin/AERserver" ]