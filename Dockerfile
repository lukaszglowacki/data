FROM golang:latest as build

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/lukaszglowacki/data
WORKDIR /go/src/github.com/lukaszglowacki/data

# Install official GoLang dep pkg
RUN go get -u github.com/golang/dep/... 

# This will install any missing packages that are listed in Gopkg
RUN dep ensure

# Build the application
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o /app/service ./cmd/service/main.go


