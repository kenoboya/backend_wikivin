FROM golang:1.22.3-alpine AS builder

WORKDIR /go/src/github.com/kenoboya/backend_wikivin

COPY . .

RUN go mod download
RUN GOOS=linux go build -o ./app ./cmd/app/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /go/src/github.com/kenoboya/backend_wikivin/app .
COPY .env .

COPY configs /configs

CMD ["./app"]
