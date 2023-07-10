FROM golang:1.20-buster

WORKDIR /go/src/app
COPY ./livetickgo .

RUN apt update
RUN apt install -y protobuf-compiler

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

RUN go mod download
RUN go generate ./...
RUN go install

CMD ["livetickgo"]

EXPOSE 9111
