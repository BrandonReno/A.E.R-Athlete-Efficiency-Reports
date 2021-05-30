# Start from golang base image
FROM golang:alpine as builder

# Add Maintainer info
LABEL maintainer="Brandon Reno <b_reno@u.pacific.edu>"

# Install git.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container 
RUN mkdir /aer

WORKDIR /aer

# Copy go mod and sum files to the working directory
COPY go.mod go.sum ./

# Download all dependencies using go mod
RUN go mod download 

#copy everything from the current working directory into the containers working directory
COPY . .

# Build the Go app as 'main.cgo'
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .






# Create a new image from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the prebuilt main executable from the previous image to the new images working dir
COPY --from=builder /aer/main .    

# Create a directory to hold the volumes
RUN mkdir ./volumes/

# Expose port 9090 for documentation
EXPOSE 9090

# Command to run the executable when containerized
ENTRYPOINT ["./main"]