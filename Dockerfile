FROM golang:1.11.5

WORKDIR /go/src/stringtheory
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build

ENTRYPOINT ["./stringtheory"]