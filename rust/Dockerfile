FROM rust:1.66.0 as builder

WORKDIR /app

COPY Cargo.lock Cargo.lock
COPY Cargo.toml Cargo.toml
COPY src/ src/

RUN cargo install --path .

FROM debian:buster-slim

RUN apt-get update && rm -rf /var/lib/apt/lists/*

COPY --from=builder /usr/local/cargo/bin/rust /usr/local/bin/udp-server

CMD ["udp-server"]
