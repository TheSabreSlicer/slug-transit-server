FROM golang:1.11

WORKDIR /go/src/sts
COPY src/ .

# RUN go get -d -v ./...
# RUN go install -v ./...
