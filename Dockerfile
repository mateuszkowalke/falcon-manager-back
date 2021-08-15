FROM golang:1.16 AS builder

ARG VERSION=dev

WORKDIR /go/src/app
COPY . .
RUN go mod tidy
RUN go build -o main -ldflags=-X=main.version=${VERSION} cmd/gql-server/main.go 

FROM debian:buster-slim
COPY --from=builder /go/src/app/main /go/bin/main
ENV PATH="/go/bin:${PATH}"
CMD ["main"]
