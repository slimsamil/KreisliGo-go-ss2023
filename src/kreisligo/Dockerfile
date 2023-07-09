FROM golang:1.20-buster
# non-go modules dependencies
RUN apt update
RUN apt install -y protobuf-compiler

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

# copy code and protobuf
WORKDIR /go/src/app
COPY ./kreisligo .
# COPY ./banktransfer/grpc/banktransfer/banktransfer.proto ./client/banktransfer/

RUN go mod download
RUN go generate ./...
RUN go install

RUN wget https://raw.githubusercontent.com/vishnubob/wait-for-it/81b1373f17855a4dc21156cfe1694c31d7d1792e/wait-for-it.sh
RUN chmod +x ./wait-for-it.sh ./docker-entrypoint.sh

ENTRYPOINT ["./docker-entrypoint.sh"]
CMD ["kreisligo"]

EXPOSE 8000
