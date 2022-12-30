build:
	go build -o bin/udp-server cmd/udp-server/main.go

serve: build
	bin/udp-server

run-dev:
	go run cmd/udp-server/main.go

docker-build:
	DOCKER_BUILDKIT=1 docker build -t go-udp-server:latest .

docker-run:
	docker run --rm -d -p 3000:3000 -p 3000:3000/udp --name go-udp-server go-udp-server:latest
