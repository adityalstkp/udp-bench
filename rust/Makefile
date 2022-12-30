docker-build:
	DOCKER_BUILDKIT=1 docker build -t rust-udp-server:latest .

docker-run:
	docker run -d --rm -p 3001:3001 -p 3001:3001/udp --name rust-udp-server rust-udp-server:latest
