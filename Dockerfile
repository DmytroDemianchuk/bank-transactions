FROM golang:1.20 AS builder

RUN go version

COPY . /github.com/dmytrodemianchuk/bank-transactions/
WORKDIR /github.com/dmytrodemianchuk/bank-transactions/

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /github.com/dmytrodemianchuk/bank-transactions/.bin/app .

CMD ["./app"]