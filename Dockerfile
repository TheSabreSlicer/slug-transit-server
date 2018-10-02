FROM golang:1.11-stretch

WORKDIR /go/src/sts
COPY src/ .

RUN go get -v ./...
RUN go run main.go
# RUN go install -v ./...

# RUN go run main.go

EXPOSE 8080

# CMD ["sts"]
