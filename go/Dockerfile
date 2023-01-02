FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY cmd/ cmd/
COPY internal/ internal/
COPY go.mod go.mod

RUN go build -o udp-server cmd/udp-server/main.go

FROM alpine:latest

RUN apk update && apk --no-cache add bash ca-certificates

COPY --from=builder /app/udp-server /udp-server

EXPOSE 3000

CMD ["/udp-server"]

