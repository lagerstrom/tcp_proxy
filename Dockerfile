FROM golang:1.23 AS builder
WORKDIR /go/src/app/

COPY go.mod go.sum cmd/main.go .


RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

FROM scratch

COPY --from=0 /go/src/app/app .

ENTRYPOINT ["/app"]
