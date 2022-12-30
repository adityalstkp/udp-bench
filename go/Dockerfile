FROM golang:1.19 AS builder

WORKDIR /app

COPY cmd/ cmd/
COPY internal/ internal/
COPY go.mod go.mod

RUN CGO_ENABLED=0 go build -o udp-server cmd/udp-server/main.go

FROM alpine

RUN apk update && apk --no-cache add bash ca-certificates

COPY --from=builder /app/udp-server /udp-server

EXPOSE 3000

CMD ["/udp-server"]

